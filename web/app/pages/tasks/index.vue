<script setup lang="ts">
import type { Task, TaskCandidate, TaskLog } from '~/types/task'

useSeoMeta({
  title: '任务中心',
  description: '待办、已办与认领任务'
})

const toast = useToast()
const { user } = useAuth()
const { queryTasks, queryClaimableTasks, claimTask, completeTask, skipTask, listTaskCandidates, queryTaskLogs } = useTaskApi()

const loadingMine = ref(false)
const loadingClaimable = ref(false)
const actingTaskId = ref<number | null>(null)
const batchClaiming = ref(false)
const mineStatus = ref<'running' | 'completed' | 'skipped'>('running')
const mineStatusOptions = [
  { label: '进行中', value: 'running' },
  { label: '已完成', value: 'completed' },
  { label: '已跳过', value: 'skipped' }
] as const

const mineTasks = ref<Task[]>([])
const claimableTasks = ref<Task[]>([])
const selectedClaimableIds = ref<number[]>([])

const detailOpen = ref(false)
const detailTask = ref<Task | null>(null)
const detailCandidates = ref<TaskCandidate[]>([])
const detailLogs = ref<TaskLog[]>([])
const detailLoadingCandidates = ref(false)
const detailLoadingLogs = ref(false)
const actionType = ref<'complete' | 'skip' | null>(null)
const actionRemark = ref('')

const statusColor = (status: Task['status']) => {
  if (status === 'running') return 'info'
  if (status === 'completed') return 'success'
  if (status === 'skipped') return 'warning'
  return 'neutral'
}

const statusText = (status: Task['status']) => {
  if (status === 'running') return '进行中'
  if (status === 'completed') return '已完成'
  if (status === 'skipped') return '已跳过'
  return '待认领'
}

const logActionText = (action: string) => {
  if (action === 'create_task') return '创建任务'
  if (action === 'claim_task') return '认领任务'
  if (action === 'complete_task') return '完成任务'
  if (action === 'skip_task') return '跳过任务'
  if (action === 'delegate_task') return '委派任务'
  return action
}

const logActionColor = (action: string) => {
  if (action === 'create_task') return 'neutral'
  if (action === 'claim_task') return 'info'
  if (action === 'complete_task') return 'success'
  if (action === 'skip_task') return 'warning'
  if (action === 'delegate_task') return 'primary'
  return 'neutral'
}

const selectedCount = computed(() => selectedClaimableIds.value.length)
const allClaimableSelected = computed(() => claimableTasks.value.length > 0 && selectedClaimableIds.value.length === claimableTasks.value.length)

const loadMineTasks = async () => {
  if (!user.value?.id) return
  loadingMine.value = true
  try {
    const result = await queryTasks({
      page: 1,
      size: 50,
      assignee_id: user.value.id,
      status: mineStatus.value
    })
    mineTasks.value = result.items || []
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载我的任务失败',
      color: 'error'
    })
  } finally {
    loadingMine.value = false
  }
}

const loadClaimableTasks = async () => {
  loadingClaimable.value = true
  try {
    const result = await queryClaimableTasks({
      page: 1,
      size: 50,
      status: 'pending'
    })
    claimableTasks.value = result.items || []
    const validIds = new Set(claimableTasks.value.map(item => item.id))
    selectedClaimableIds.value = selectedClaimableIds.value.filter(id => validIds.has(id))
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载可认领任务失败',
      color: 'error'
    })
  } finally {
    loadingClaimable.value = false
  }
}

const refreshAll = async () => {
  await Promise.all([loadMineTasks(), loadClaimableTasks()])
}

const toggleClaimable = (taskId: number, checked: boolean) => {
  if (checked) {
    if (!selectedClaimableIds.value.includes(taskId)) {
      selectedClaimableIds.value = [...selectedClaimableIds.value, taskId]
    }
    return
  }
  selectedClaimableIds.value = selectedClaimableIds.value.filter(id => id !== taskId)
}

const toggleAllClaimable = (checked: boolean) => {
  selectedClaimableIds.value = checked ? claimableTasks.value.map(item => item.id) : []
}

const onClaim = async (taskId: number) => {
  actingTaskId.value = taskId
  try {
    await claimTask({ id: taskId })
    toast.add({
      title: '认领成功',
      description: '任务已进入你的待办',
      color: 'success'
    })
    await refreshAll()
  } catch (err: unknown) {
    toast.add({
      title: '认领失败',
      description: err instanceof Error ? err.message : '认领任务失败',
      color: 'error'
    })
  } finally {
    actingTaskId.value = null
  }
}

const onBatchClaim = async () => {
  if (selectedClaimableIds.value.length === 0) return
  batchClaiming.value = true
  try {
    for (const taskId of selectedClaimableIds.value) {
      await claimTask({ id: taskId })
    }
    toast.add({
      title: '批量认领成功',
      description: `已认领 ${selectedClaimableIds.value.length} 个任务`,
      color: 'success'
    })
    selectedClaimableIds.value = []
    await refreshAll()
  } catch (err: unknown) {
    toast.add({
      title: '批量认领失败',
      description: err instanceof Error ? err.message : '请稍后重试',
      color: 'error'
    })
  } finally {
    batchClaiming.value = false
  }
}

const loadTaskDetailExtras = async (taskId: number) => {
  detailLoadingCandidates.value = true
  detailLoadingLogs.value = true
  const [candidateResult, logResult] = await Promise.allSettled([
    listTaskCandidates({ task_id: taskId }),
    queryTaskLogs({ page: 1, size: 100, task_id: taskId })
  ])

  if (candidateResult.status === 'fulfilled') {
    detailCandidates.value = candidateResult.value || []
  } else {
    detailCandidates.value = []
    toast.add({
      title: '候选人加载失败',
      description: candidateResult.reason instanceof Error ? candidateResult.reason.message : '查询任务候选人失败',
      color: 'error'
    })
  }

  if (logResult.status === 'fulfilled') {
    detailLogs.value = logResult.value.items || []
  } else {
    detailLogs.value = []
    toast.add({
      title: '日志加载失败',
      description: logResult.reason instanceof Error ? logResult.reason.message : '查询任务日志失败',
      color: 'error'
    })
  }
  detailLoadingCandidates.value = false
  detailLoadingLogs.value = false
}

const openTaskDetail = async (task: Task) => {
  detailTask.value = task
  detailCandidates.value = []
  detailLogs.value = []
  actionType.value = null
  actionRemark.value = task.remark || ''
  detailOpen.value = true
  await loadTaskDetailExtras(task.id)
}

const beginAction = (type: 'complete' | 'skip') => {
  actionType.value = type
}

const submitAction = async () => {
  if (!detailTask.value || !user.value?.id || !actionType.value) return
  actingTaskId.value = detailTask.value.id
  try {
    if (actionType.value === 'complete') {
      await completeTask({
        id: detailTask.value.id,
        assignee_id: user.value.id,
        remark: actionRemark.value
      })
    } else {
      await skipTask({
        id: detailTask.value.id,
        assignee_id: user.value.id
      })
    }
    toast.add({
      title: '操作成功',
      description: actionType.value === 'complete' ? '任务已完成' : '任务已跳过',
      color: 'success'
    })
    detailOpen.value = false
    await refreshAll()
  } catch (err: unknown) {
    toast.add({
      title: '操作失败',
      description: err instanceof Error ? err.message : '请稍后重试',
      color: 'error'
    })
  } finally {
    actingTaskId.value = null
    actionType.value = null
  }
}

watch(mineStatus, () => {
  loadMineTasks()
})

watch(detailOpen, (opened) => {
  if (opened) return
  detailTask.value = null
  detailCandidates.value = []
  detailLogs.value = []
  actionType.value = null
  actionRemark.value = ''
})

onMounted(() => {
  refreshAll()
})
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="任务中心"
      description="管理待办、已办与认领池任务"
    />

    <div class="flex items-center justify-end">
      <UButton
        icon="i-lucide-refresh-cw"
        color="neutral"
        variant="soft"
        :loading="loadingMine || loadingClaimable"
        @click="refreshAll"
      >
        刷新
      </UButton>
    </div>

    <div class="grid gap-4 lg:grid-cols-2">
      <UCard>
        <template #header>
          <div class="flex items-center justify-between gap-2">
            <h3 class="font-semibold">
              我的任务
            </h3>
            <USelect
              v-model="mineStatus"
              class="w-36"
              :items="mineStatusOptions"
              value-key="value"
              label-key="label"
            />
          </div>
        </template>

        <div
          v-if="loadingMine"
          class="py-8 text-sm text-muted"
        >
          正在加载...
        </div>
        <div
          v-else-if="mineTasks.length === 0"
          class="py-8 text-sm text-muted"
        >
          暂无任务
        </div>
        <div
          v-else
          class="space-y-3"
        >
          <div
            v-for="task in mineTasks"
            :key="task.id"
            class="rounded-lg border border-default p-3"
          >
            <div class="flex items-start justify-between gap-2">
              <div class="space-y-1">
                <p class="text-sm font-medium">
                  {{ task.node_name }}
                </p>
                <p class="text-xs text-muted">
                  {{ task.process_name }} · 实例 #{{ task.execution_id }}
                </p>
              </div>
              <UBadge
                :color="statusColor(task.status)"
                variant="soft"
              >
                {{ statusText(task.status) }}
              </UBadge>
            </div>

            <div class="mt-3 flex items-center gap-2">
              <UButton
                size="xs"
                color="neutral"
                variant="soft"
                @click="openTaskDetail(task)"
              >
                详情
              </UButton>
            </div>
          </div>
        </div>
      </UCard>

      <UCard>
        <template #header>
          <div class="flex items-center justify-between gap-2">
            <h3 class="font-semibold">
              可认领任务
            </h3>
            <div class="flex items-center gap-2">
              <label class="flex items-center gap-1 text-xs text-muted">
                <input
                  type="checkbox"
                  :checked="allClaimableSelected"
                  @change="toggleAllClaimable(($event.target as HTMLInputElement).checked)"
                >
                全选
              </label>
              <UButton
                size="xs"
                color="primary"
                variant="soft"
                :loading="batchClaiming"
                :disabled="selectedCount === 0"
                @click="onBatchClaim"
              >
                批量认领({{ selectedCount }})
              </UButton>
            </div>
          </div>
        </template>

        <div
          v-if="loadingClaimable"
          class="py-8 text-sm text-muted"
        >
          正在加载...
        </div>
        <div
          v-else-if="claimableTasks.length === 0"
          class="py-8 text-sm text-muted"
        >
          暂无可认领任务
        </div>
        <div
          v-else
          class="space-y-3"
        >
          <div
            v-for="task in claimableTasks"
            :key="task.id"
            class="rounded-lg border border-default p-3"
          >
            <div class="flex items-start justify-between gap-2">
              <div class="flex items-start gap-2">
                <input
                  class="mt-1"
                  type="checkbox"
                  :checked="selectedClaimableIds.includes(task.id)"
                  @change="toggleClaimable(task.id, ($event.target as HTMLInputElement).checked)"
                >
                <div class="space-y-1">
                  <p class="text-sm font-medium">
                    {{ task.node_name }}
                  </p>
                  <p class="text-xs text-muted">
                    {{ task.process_name }} · 实例 #{{ task.execution_id }}
                  </p>
                </div>
              </div>
              <UBadge
                :color="statusColor(task.status)"
                variant="soft"
              >
                {{ statusText(task.status) }}
              </UBadge>
            </div>
            <div class="mt-3 flex items-center gap-2">
              <UButton
                size="xs"
                color="neutral"
                variant="soft"
                @click="openTaskDetail(task)"
              >
                详情
              </UButton>
              <UButton
                size="xs"
                color="primary"
                variant="soft"
                :loading="actingTaskId === task.id"
                @click="onClaim(task.id)"
              >
                认领
              </UButton>
            </div>
          </div>
        </div>
      </UCard>
    </div>

    <UModal v-model:open="detailOpen">
      <template #content>
        <UCard>
          <template #header>
            <div class="flex items-center justify-between">
              <h3 class="font-semibold">
                任务详情
              </h3>
              <UBadge
                v-if="detailTask"
                :color="statusColor(detailTask.status)"
                variant="soft"
              >
                {{ statusText(detailTask.status) }}
              </UBadge>
            </div>
          </template>

          <div
            v-if="detailTask"
            class="max-h-[70vh] space-y-4 overflow-y-auto text-sm"
          >
            <div class="grid gap-2 rounded-lg border border-default p-3">
              <p><span class="text-muted">任务ID:</span> {{ detailTask.id }}</p>
              <p><span class="text-muted">流程:</span> {{ detailTask.process_name }}</p>
              <p><span class="text-muted">节点:</span> {{ detailTask.node_name }}</p>
              <p><span class="text-muted">实例ID:</span> {{ detailTask.execution_id }}</p>
              <p><span class="text-muted">开始时间:</span> {{ detailTask.started_at || '-' }}</p>
              <p><span class="text-muted">认领时间:</span> {{ detailTask.claimed_at || '-' }}</p>
              <p><span class="text-muted">结束时间:</span> {{ detailTask.ended_at || '-' }}</p>
            </div>

            <div class="space-y-2 rounded-lg border border-default p-3">
              <p class="font-medium">
                候选人
              </p>
              <div
                v-if="detailLoadingCandidates"
                class="text-xs text-muted"
              >
                正在加载候选人...
              </div>
              <div
                v-else-if="detailCandidates.length === 0"
                class="text-xs text-muted"
              >
                无候选人（可能为直接指派任务）
              </div>
              <div
                v-else
                class="flex flex-wrap gap-2"
              >
                <UBadge
                  v-for="item in detailCandidates"
                  :key="item.id"
                  color="neutral"
                  variant="soft"
                >
                  {{ item.user_name || `用户#${item.user_id}` }}
                </UBadge>
              </div>
            </div>

            <div class="space-y-2 rounded-lg border border-default p-3">
              <p class="font-medium">
                操作日志
              </p>
              <div
                v-if="detailLoadingLogs"
                class="text-xs text-muted"
              >
                正在加载日志...
              </div>
              <div
                v-else-if="detailLogs.length === 0"
                class="text-xs text-muted"
              >
                暂无日志
              </div>
              <div
                v-else
                class="space-y-2"
              >
                <div
                  v-for="log in detailLogs"
                  :key="log.id"
                  class="rounded-md border border-default p-2"
                >
                  <div class="flex items-center gap-2">
                    <UBadge
                      :color="logActionColor(log.action)"
                      variant="soft"
                    >
                      {{ logActionText(log.action) }}
                    </UBadge>
                    <span class="text-xs text-muted">{{ log.created_at }}</span>
                  </div>
                  <p
                    v-if="log.assignee_name"
                    class="mt-1 text-xs text-muted"
                  >
                    执行人: {{ log.assignee_name }}
                  </p>
                  <p
                    v-if="log.remark"
                    class="mt-1 text-xs"
                  >
                    备注: {{ log.remark }}
                  </p>
                </div>
              </div>
            </div>

            <UFormField
              v-if="detailTask.status === 'running'"
              label="处理备注"
            >
              <UTextarea
                v-model="actionRemark"
                class="w-full"
                :rows="3"
                placeholder="可选，填写处理说明"
              />
            </UFormField>
          </div>

          <template #footer>
            <div class="flex justify-end gap-2">
              <UButton
                color="neutral"
                variant="ghost"
                @click="detailOpen = false"
              >
                关闭
              </UButton>
              <template v-if="detailTask?.status === 'running'">
                <UButton
                  color="warning"
                  variant="soft"
                  :loading="actingTaskId === detailTask?.id && actionType === 'skip'"
                  @click="beginAction('skip')"
                >
                  跳过
                </UButton>
                <UButton
                  color="success"
                  :loading="actingTaskId === detailTask?.id && actionType === 'complete'"
                  @click="beginAction('complete')"
                >
                  完成
                </UButton>
                <UButton
                  v-if="actionType"
                  color="primary"
                  variant="soft"
                  :loading="actingTaskId === detailTask?.id"
                  @click="submitAction"
                >
                  确认{{ actionType === 'complete' ? '完成' : '跳过' }}
                </UButton>
              </template>
            </div>
          </template>
        </UCard>
      </template>
    </UModal>
  </UContainer>
</template>
