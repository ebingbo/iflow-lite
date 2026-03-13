<script setup lang="ts">
import type { TableColumn } from '#ui/components/Table.vue'
import type { Execution } from '~/types/execution'
import type { Task, TaskLog } from '~/types/task'
import type { Row } from '@tanstack/vue-table'
import type { DropdownMenuItem } from '#ui/components/DropdownMenu.vue'
import dayjs from 'dayjs'

useSeoMeta({
  title: '流程运行',
  description: '流程实例与运行轨迹'
})

const toast = useToast()
const { queryExecutions, getExecution } = useExecutionApi()
const { queryTasks, queryTaskLogs } = useTaskApi()

const loading = ref(false)
const listRequestId = ref(0)
const executions = ref<Execution[]>([])
const total = ref(0)
const page = ref(1)
const size = ref(20)

const status = ref<'running' | 'completed' | ''>('')
const processCode = ref('')
const businessKey = ref('')

const detailOpen = ref(false)
const detailExecution = ref<Execution | null>(null)
const detailLogs = ref<TaskLog[]>([])
const detailTasks = ref<Task[]>([])
const detailLoading = ref(false)

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')

const columns: TableColumn<Execution>[] = [
  {
    accessorKey: 'id',
    header: '#',
    cell: ({ row }) => `#${row.getValue('id')}`
  },
  {
    accessorKey: 'process_name',
    header: '流程名称'
  },
  {
    accessorKey: 'process_code',
    header: '流程编码'
  },
  {
    accessorKey: 'business_key',
    header: '业务标识'
  },
  {
    accessorKey: 'status',
    header: '状态',
    cell: ({ row }) => {
      const value = row.getValue('status') as string
      return h(UBadge, { variant: 'soft', color: statusColor(value) }, () => statusText(value))
    }
  },
  {
    accessorKey: 'started_at',
    header: '开始时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('created_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    accessorKey: 'ended_at',
    header: '结束时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('ended_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    accessorKey: 'progress',
    header: '进度',
    cell: ({ row }) => `${row.getValue('progress') ?? 0}%`
  },
  {
    id: 'actions',
    header: '操作',
    meta: {
      class: {
        td: 'text-right'
      }
    },
    cell: ({ row }) => {
      return h(
        UDropdownMenu,
        {
          items: getRowItems(row)
        },
        () =>
          h(UButton, {
            'icon': 'i-lucide-ellipsis-vertical',
            'color': 'neutral',
            'variant': 'ghost',
            'aria-label': 'Actions dropdown'
          })
      )
    }
  }
]

function getRowItems(row: Row<Execution>) {
  const items: DropdownMenuItem[] = [
    {
      label: '详情',
      onSelect() {
        openDetail(row.original)
      }
    }
  ]
  return items
}

const statusColor = (value: string) => {
  if (value === 'running') return 'info'
  if (value === 'completed') return 'success'
  return 'neutral'
}

const statusText = (value: string) => {
  if (value === 'running') return '进行中'
  if (value === 'completed') return '已完成'
  return value || '-'
}

const logActionText = (action: string) => {
  if (action === 'start_process') return '启动流程'
  if (action === 'complete_process') return '完成流程'
  if (action === 'create_task') return '创建任务'
  if (action === 'claim_task') return '认领任务'
  if (action === 'complete_task') return '完成任务'
  if (action === 'skip_task') return '跳过任务'
  if (action === 'delegate_task') return '委派任务'
  if (action === 'join_task') return '汇聚任务'
  return action
}

const loadExecutions = async () => {
  const requestId = ++listRequestId.value
  const timeoutMs = 8000
  let timeoutId: ReturnType<typeof setTimeout> | undefined
  const timeoutPromise = new Promise<never>((_, reject) => {
    timeoutId = setTimeout(() => reject(new Error('请求超时，请重试')), timeoutMs)
  })

  try {
    const result = await Promise.race([
      queryExecutions({
        page: page.value,
        size: size.value,
        process_code: processCode.value || undefined,
        business_key: businessKey.value || undefined,
        status: status.value || undefined
      }),
      timeoutPromise
    ])
    executions.value = result.items || []
    total.value = result.total || 0
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载实例失败',
      color: 'error'
    })
  } finally {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }
    if (requestId === listRequestId.value) {
      loading.value = false
    }
  }
}

const refreshList = async () => {
  page.value = 1
  await loadExecutions()
}

const currentTasks = computed(() => detailTasks.value.filter(item => item.status === 'running' || item.status === 'pending'))
const nodeNameMap = computed(() => {
  const map = new Map<number, string>()
  for (const task of detailTasks.value) {
    map.set(task.node_id, task.node_name)
  }
  return map
})

const logItems = computed(() => {
  const items = [...detailLogs.value]
  items.sort((a, b) => {
    if (!a.created_at || !b.created_at) return 0
    return a.created_at.localeCompare(b.created_at)
  })
  return items
})

const openDetail = async (execution: Execution) => {
  detailExecution.value = execution
  detailOpen.value = true
  detailLoading.value = true
  try {
    const [executionResult, taskResult, logResult] = await Promise.all([
      getExecution(execution.id),
      queryTasks({
        page: 1,
        size: 200,
        execution_id: execution.id
      }),
      queryTaskLogs({
        page: 1,
        size: 200,
        execution_id: execution.id
      })
    ])
    detailExecution.value = executionResult || execution
    detailTasks.value = taskResult.items || []
    detailLogs.value = logResult.items || []
  } catch (err: unknown) {
    toast.add({
      title: '加载详情失败',
      description: err instanceof Error ? err.message : '加载实例详情失败',
      color: 'error'
    })
  } finally {
    detailLoading.value = false
  }
}

watch([status, processCode, businessKey], () => {
  refreshList()
})

watch(detailOpen, (opened) => {
  if (opened) return
  detailExecution.value = null
  detailLogs.value = []
  detailTasks.value = []
})

onMounted(() => {
  loadExecutions()
})
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="流程运行"
      description="查看流程实例、状态、运行轨迹"
    />

    <UCard>
      <div class="flex flex-wrap items-end gap-3">
        <UFormField label="流程编码">
          <UInput
            v-model="processCode"
            placeholder="输入流程编码"
          />
        </UFormField>
        <UFormField label="业务标识">
          <UInput
            v-model="businessKey"
            placeholder="输入业务Key"
          />
        </UFormField>
        <UFormField label="状态">
          <USelect
            v-model="status"
            class="w-32"
            :items="[
              { label: '进行中', value: 'running' },
              { label: '已完成', value: 'completed' }
            ]"
            placeholder="选择状态"
          />
        </UFormField>
        <UButton
          color="neutral"
          variant="soft"
          :loading="loading"
          @click="refreshList"
        >
          刷新
        </UButton>
      </div>
    </UCard>

    <UCard>
      <div
        v-if="loading && executions.length === 0"
        class="py-8 text-sm text-muted"
      >
        正在加载...
      </div>
      <div
        v-else-if="executions.length === 0"
        class="py-8 text-sm text-muted"
      >
        暂无实例
      </div>
      <div
        v-else
        class="space-y-3"
      >
        <UTable
          :columns="columns"
          :data="executions"
          :loading="loading"
        />
      </div>
    </UCard>

    <USlideover v-model:open="detailOpen">
      <template #content>
        <UCard>
          <template #header>
            <div class="flex items-center justify-between gap-2">
              <div>
                <p class="text-sm font-semibold">
                  {{ detailExecution?.process_name || '实例详情' }}
                </p>
                <p class="text-xs text-muted">
                  实例 #{{ detailExecution?.id || '-' }}
                </p>
              </div>
              <UBadge
                v-if="detailExecution"
                :color="statusColor(detailExecution.status)"
                variant="soft"
              >
                {{ statusText(detailExecution.status) }}
              </UBadge>
            </div>
          </template>

          <div class="max-h-[75vh] space-y-4 overflow-y-auto text-sm">
            <div class="grid gap-2 rounded-lg border border-default p-3">
              <p><span class="text-muted">流程编码:</span> {{ detailExecution?.process_code || '-' }}</p>
              <p><span class="text-muted">业务标识:</span> {{ detailExecution?.business_key || '-' }}</p>
              <p><span class="text-muted">创建人:</span> {{ detailExecution?.created_by || '-' }}</p>
              <p><span class="text-muted">开始时间:</span> {{ detailExecution?.started_at || '-' }}</p>
              <p><span class="text-muted">结束时间:</span> {{ detailExecution?.ended_at || '-' }}</p>
            </div>

            <div class="space-y-2 rounded-lg border border-default p-3">
              <p class="font-medium">
                当前卡点
              </p>
              <div
                v-if="detailLoading"
                class="text-xs text-muted"
              >
                正在加载...
              </div>
              <div
                v-else-if="currentTasks.length === 0"
                class="text-xs text-muted"
              >
                暂无未完成节点
              </div>
              <div
                v-else
                class="space-y-2"
              >
                <div
                  v-for="task in currentTasks"
                  :key="task.id"
                  class="rounded-md border border-default p-2"
                >
                  <div class="flex items-center justify-between gap-2">
                    <p class="text-sm font-medium">
                      {{ task.node_name }}
                    </p>
                    <UBadge
                      :color="statusColor(task.status)"
                      variant="soft"
                    >
                      {{ statusText(task.status) }}
                    </UBadge>
                  </div>
                  <p class="mt-1 text-xs text-muted">
                    处理人: {{ task.assignee_name || (task.assignee_id ? `用户#${task.assignee_id}` : '-') }}
                  </p>
                </div>
              </div>
            </div>

            <div class="space-y-2 rounded-lg border border-default p-3">
              <p class="font-medium">
                节点轨迹
              </p>
              <div
                v-if="detailLoading"
                class="text-xs text-muted"
              >
                正在加载...
              </div>
              <div
                v-else-if="logItems.length === 0"
                class="text-xs text-muted"
              >
                暂无轨迹
              </div>
              <div
                v-else
                class="space-y-2"
              >
                <div
                  v-for="log in logItems"
                  :key="log.id"
                  class="rounded-md border border-default p-2"
                >
                  <div class="flex items-center justify-between gap-2">
                    <p class="text-sm font-medium">
                      {{ nodeNameMap.get(log.node_id) || log.node_code || `节点#${log.node_id}` }}
                    </p>
                    <span class="text-xs text-muted">
                      {{ log.created_at }}
                    </span>
                  </div>
                  <p class="mt-1 text-xs text-muted">
                    动作: {{ logActionText(log.action) }}
                  </p>
                  <p
                    v-if="log.assignee_name"
                    class="mt-1 text-xs text-muted"
                  >
                    处理人: {{ log.assignee_name }}
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
          </div>

          <template #footer>
            <div class="flex justify-end">
              <UButton
                color="neutral"
                variant="ghost"
                @click="detailOpen = false"
              >
                关闭
              </UButton>
            </div>
          </template>
        </UCard>
      </template>
    </USlideover>
  </UContainer>
</template>
