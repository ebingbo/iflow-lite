import type { ApiResponse } from '~/types/api'
import type { RoleListBody, RoleOption } from '~/types/user'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useRoleApi() {
  async function listRoles(body: RoleListBody = {}) {
    const response = await $fetch<ApiResponse<RoleOption[]>>('/api/role/list', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取角色列表失败')
  }

  return {
    listRoles
  }
}
