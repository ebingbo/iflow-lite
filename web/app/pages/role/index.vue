<script setup lang="ts">
import type { TableColumn } from '#ui/components/Table.vue'
import type { Role } from '~/types/role'
import type { Row } from '@tanstack/vue-table'
import type { DropdownMenuItem } from '#ui/components/DropdownMenu.vue'
import dayjs from 'dayjs'

useSeoMeta({
  title: '角色管理',
  description: '系统角色管理'
})

const toast = useToast()
const { queryRoles, addRole, updateRole, deleteRole } = useRoleApi()

const roles = ref<Role[]>([])
const loading = ref(false)
const keyword = ref('')

const editingRole = ref<Role | null>(null)
const editOpen = ref(false)
const formName = ref('')
const formCode = ref('')
const saving = ref(false)

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')

const columns: TableColumn<Role>[] = [
  {
    accessorKey: 'id',
    header: '#',
    cell: ({ row }) => `#${row.getValue('id')}`
  },
  {
    accessorKey: 'name',
    header: '名称'
  },
  {
    accessorKey: 'code',
    header: '编码'
  },
  {
    accessorKey: 'created_at',
    header: '创建时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('created_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    accessorKey: 'updated_at',
    header: '更新时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('updated_at')).format('YYYY-MM-DD HH:mm:ss')
    }
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

function getRowItems(row: Row<Role>) {
  const items: DropdownMenuItem[] = [
    {
      label: '编辑',
      onSelect() {
        openEdit(row.original)
      }
    },
    {
      label: '删除',
      async onSelect() {
        await removeRole(row.original)
      }
    }
  ]
  return items
}

const loadRoles = async () => {
  loading.value = true
  try {
    const result = await queryRoles({
      page: 1,
      size: 50,
      keyword: keyword.value || undefined
    })
    roles.value = result.items || []
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载角色失败',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

const refresh = async () => {
  await loadRoles()
}

const openCreate = () => {
  editingRole.value = null
  formName.value = ''
  formCode.value = ''
  editOpen.value = true
}

const openEdit = (role: Role) => {
  editingRole.value = role
  formName.value = role.name
  formCode.value = role.code
  editOpen.value = true
}

const saveRole = async () => {
  if (!formName.value || !formCode.value) {
    toast.add({
      title: '提示',
      description: '名称和编码不能为空',
      color: 'warning'
    })
    return
  }
  saving.value = true
  try {
    if (editingRole.value) {
      await updateRole({
        id: editingRole.value.id,
        name: formName.value,
        code: formCode.value
      })
    } else {
      await addRole({
        name: formName.value,
        code: formCode.value
      })
    }
    toast.add({
      title: '保存成功',
      color: 'success'
    })
    editOpen.value = false
    await loadRoles()
  } catch (err: unknown) {
    toast.add({
      title: '保存失败',
      description: err instanceof Error ? err.message : '保存角色失败',
      color: 'error'
    })
  } finally {
    saving.value = false
  }
}

const removeRole = async (role: Role) => {
  if (!confirm(`确定删除角色「${role.name}」吗？`)) return
  try {
    await deleteRole(role.id)
    toast.add({
      title: '删除成功',
      color: 'success'
    })
    await loadRoles()
  } catch (err: unknown) {
    toast.add({
      title: '删除失败',
      description: err instanceof Error ? err.message : '删除角色失败',
      color: 'error'
    })
  }
}

watch(keyword, () => {
  refresh()
})

watch(editOpen, (opened) => {
  if (opened) return
  editingRole.value = null
  formName.value = ''
  formCode.value = ''
})

onMounted(() => {
  loadRoles()
})
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="角色管理"
      description="管理角色及分派权限"
    />

    <UCard>
      <div class="flex flex-wrap items-end gap-3">
        <UFormField label="关键词">
          <UInput
            v-model="keyword"
            placeholder="名称或编码"
          />
        </UFormField>
        <UButton
          color="primary"
          @click="openCreate"
        >
          新增角色
        </UButton>
        <UButton
          color="neutral"
          variant="soft"
          :loading="loading"
          @click="refresh"
        >
          刷新
        </UButton>
      </div>
    </UCard>

    <UCard>
      <UTable
        :columns="columns"
        :data="roles"
        :loading="loading"
      />
      <div
        v-if="!loading && roles.length === 0"
        class="py-6 text-sm text-muted"
      >
        暂无角色
      </div>
    </UCard>

    <UModal v-model:open="editOpen">
      <template #content>
        <UCard>
          <template #header>
            <p class="text-sm font-semibold">
              {{ editingRole ? '编辑角色' : '新增角色' }}
            </p>
          </template>

          <div class="space-y-3">
            <UFormField label="角色名称">
              <UInput
                v-model="formName"
                placeholder="输入角色名称"
              />
            </UFormField>
            <UFormField label="角色编码">
              <UInput
                v-model="formCode"
                placeholder="输入角色编码"
              />
            </UFormField>
          </div>

          <template #footer>
            <div class="flex justify-end gap-2">
              <UButton
                color="neutral"
                variant="ghost"
                @click="editOpen = false"
              >
                关闭
              </UButton>
              <UButton
                color="primary"
                :loading="saving"
                @click="saveRole"
              >
                保存
              </UButton>
            </div>
          </template>
        </UCard>
      </template>
    </UModal>
  </UContainer>
</template>
