import type { ApiResponse } from '~/types/api'
import type { UserListBody, UserOption } from '~/types/user'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useUserApi() {
  async function listUsers(body: UserListBody = {}) {
    const response = await $fetch<ApiResponse<UserOption[]>>('/api/user/list', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取用户列表失败')
  }

  return {
    listUsers
  }
}
