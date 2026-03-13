import type { ApiResponse } from '~/types/api'
import type {
  User,
  UserListBody,
  UserOption,
  UserQueryBody,
  UserQueryData,
  UserRoleListBody,
  UserRoleListData,
  UserRoleUpdateBody,
  UserStatusUpdateBody
} from '~/types/user'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useUserApi() {
  async function queryUsers(body: UserQueryBody) {
    const response = await $fetch<ApiResponse<UserQueryData>>('/api/user/query', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '查询用户失败')
  }

  async function listUsers(body: UserListBody = {}) {
    const response = await $fetch<ApiResponse<UserOption[]>>('/api/user/list', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取用户列表失败')
  }

  async function updateUserStatus(body: UserStatusUpdateBody) {
    const response = await $fetch<ApiResponse<User>>('/api/user/status/update', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '更新用户状态失败')
  }

  async function listUserRoles(body: UserRoleListBody) {
    const response = await $fetch<ApiResponse<UserRoleListData>>('/api/user/role/list', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取用户角色失败')
  }

  async function updateUserRoles(body: UserRoleUpdateBody) {
    const response = await $fetch<ApiResponse<null>>('/api/user/role/update', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '更新用户角色失败')
  }

  return {
    queryUsers,
    listUsers,
    updateUserStatus,
    listUserRoles,
    updateUserRoles
  }
}
