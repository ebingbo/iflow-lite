<script setup lang="ts">
import type { TableColumn } from '#ui/components/Table.vue'
import type { User } from '~/types/user'
import type { Role } from '~/types/role'
import type { Row } from '@tanstack/vue-table'
import type { DropdownMenuItem } from '#ui/components/DropdownMenu.vue'
import dayjs from 'dayjs'

useSeoMeta({
  title: '用户管理',
  description: '系统用户管理'
})

const toast = useToast()
const { queryUsers, updateUserStatus, listUserRoles, updateUserRoles } = useUserApi()
const { queryRoles } = useRoleApi()

const users = ref<User[]>([])
const loading = ref(false)

const page = ref(1)
const size = ref(20)
const name = ref('')
const email = ref('')
const status = ref<number | ''>('')

const roleDialogOpen = ref(false)
const roleLoading = ref(false)
const selectedUser = ref<User | null>(null)
const roleOptions = ref<Role[]>([])
const selectedRoleIds = ref<number[]>([])

const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')

const columns: TableColumn<User>[] = [
  {
    accessorKey: 'id',
    header: '#',
    cell: ({ row }) => `#${row.getValue('id')}`
  },
  {
    accessorKey: 'name',
    header: '姓名'
  },
  {
    accessorKey: 'email',
    header: '邮箱'
  },
  {
    accessorKey: 'status',
    header: '状态',
    cell: ({ row }) => {
      const value = row.getValue('status') as number | null
      const color = value === 1 ? 'success' : 'neutral'
      return h(UBadge, { variant: 'soft', color }, () => (value === 1 ? '启用' : '禁用'))
    }
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

function getRowItems(row: Row<User>) {
  const items: DropdownMenuItem[] = [
    {
      label: '角色绑定',
      onSelect() {
        openRoleDialog(row.original)
      }
    },
    {
      type: 'separator'
    }
  ]
  if ((row.original.status ?? 0) === 1) {
    items.push({
      label: '禁用',
      async onSelect() {
        await toggleUserStatus(row.original, 0)
      }
    })
  } else {
    items.push({
      label: '启用',
      async onSelect() {
        await toggleUserStatus(row.original, 1)
      }
    })
  }
  return items
}

const loadUsers = async () => {
  loading.value = true
  try {
    const result = await queryUsers({
      page: page.value,
      size: size.value,
      name: name.value || undefined,
      email: email.value || undefined,
      status: status.value === '' ? undefined : status.value
    })
    users.value = result.items || []
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载用户失败',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}

const refresh = async () => {
  page.value = 1
  await loadUsers()
}

const toggleUserStatus = async (user: User, nextStatus: number) => {
  try {
    await updateUserStatus({ id: user.id, status: nextStatus })
    toast.add({
      title: '操作成功',
      description: nextStatus === 1 ? '用户已启用' : '用户已禁用',
      color: 'success'
    })
    await loadUsers()
  } catch (err: unknown) {
    toast.add({
      title: '操作失败',
      description: err instanceof Error ? err.message : '更新用户状态失败',
      color: 'error'
    })
  }
}

const openRoleDialog = async (user: User) => {
  selectedUser.value = user
  roleDialogOpen.value = true
  roleLoading.value = true
  try {
    const [rolesResult, userRolesResult] = await Promise.all([
      queryRoles({ page: 1, size: 200 }),
      listUserRoles({ user_id: user.id })
    ])
    roleOptions.value = rolesResult.items || []
    selectedRoleIds.value = userRolesResult || []
  } catch (err: unknown) {
    toast.add({
      title: '加载失败',
      description: err instanceof Error ? err.message : '加载角色失败',
      color: 'error'
    })
  } finally {
    roleLoading.value = false
  }
}

const toggleRole = (roleId: number, checked: boolean) => {
  if (checked) {
    if (!selectedRoleIds.value.includes(roleId)) {
      selectedRoleIds.value = [...selectedRoleIds.value, roleId]
    }
    return
  }
  selectedRoleIds.value = selectedRoleIds.value.filter(id => id !== roleId)
}

const saveUserRoles = async () => {
  if (!selectedUser.value) return
  try {
    await updateUserRoles({
      user_id: selectedUser.value.id,
      role_ids: selectedRoleIds.value
    })
    toast.add({
      title: '保存成功',
      description: '角色绑定已更新',
      color: 'success'
    })
    roleDialogOpen.value = false
  } catch (err: unknown) {
    toast.add({
      title: '保存失败',
      description: err instanceof Error ? err.message : '更新角色失败',
      color: 'error'
    })
  }
}

watch([name, email, status], () => {
  refresh()
})

watch(roleDialogOpen, (opened) => {
  if (opened) return
  selectedUser.value = null
  roleOptions.value = []
  selectedRoleIds.value = []
})

onMounted(() => {
  loadUsers()
})
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="用户管理"
      description="管理用户信息与启停状态"
    />

    <UCard>
      <div class="flex flex-wrap items-end gap-3">
        <UFormField label="姓名">
          <UInput
            v-model="name"
            placeholder="输入姓名"
          />
        </UFormField>
        <UFormField label="邮箱">
          <UInput
            v-model="email"
            placeholder="输入邮箱"
          />
        </UFormField>
        <UFormField label="状态">
          <USelect
            v-model="status"
            class="w-32"
            :items="[
              { label: '启用', value: 1 },
              { label: '禁用', value: 0 }
            ]"
            value-key="value"
            label-key="label"
            placeholder="选择状态"
          />
        </UFormField>
        <UButton
          color="neutral"
          variant="soft"
          @click="refresh"
        >
          刷新
        </UButton>
      </div>
    </UCard>

    <UCard>
      <UTable
        :columns="columns"
        :data="users"
        :loading="loading"
      />
      <div
        v-if="!loading && users.length === 0"
        class="py-6 text-sm text-muted"
      >
        暂无用户
      </div>
    </UCard>

    <UModal v-model:open="roleDialogOpen">
      <template #content>
        <UCard>
          <template #header>
            <div class="flex items-center justify-between gap-2">
              <div>
                <p class="text-sm font-semibold">
                  角色绑定
                </p>
                <p class="text-xs text-muted">
                  {{ selectedUser?.name || '-' }}
                </p>
              </div>
              <UBadge
                v-if="selectedUser"
                :color="selectedUser.status === 1 ? 'success' : 'neutral'"
                variant="soft"
              >
                {{ selectedUser.status === 1 ? '启用' : '禁用' }}
              </UBadge>
            </div>
          </template>

          <div class="space-y-2 text-sm">
            <div
              v-if="roleLoading"
              class="text-xs text-muted"
            >
              正在加载角色...
            </div>
            <div
              v-else-if="roleOptions.length === 0"
              class="text-xs text-muted"
            >
              暂无可选角色
            </div>
            <div
              v-else
              class="space-y-2"
            >
              <label
                v-for="role in roleOptions"
                :key="role.id"
                class="flex items-center gap-2 rounded-md border border-default px-3 py-2"
              >
                <input
                  type="checkbox"
                  :checked="selectedRoleIds.includes(role.id)"
                  @change="toggleRole(role.id, ($event.target as HTMLInputElement).checked)"
                >
                <div class="flex-1">
                  <p class="text-sm font-medium">
                    {{ role.name }}
                  </p>
                  <p class="text-xs text-muted">
                    {{ role.code }}
                  </p>
                </div>
              </label>
            </div>
          </div>

          <template #footer>
            <div class="flex justify-end gap-2">
              <UButton
                color="neutral"
                variant="ghost"
                @click="roleDialogOpen = false"
              >
                关闭
              </UButton>
              <UButton
                color="primary"
                :loading="roleLoading"
                @click="saveUserRoles"
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
