import type { ApiResponse } from '~/types/api'
import type { RoleListBody, RoleOption } from '~/types/user'
import type { Role, RoleAddBody, RoleQueryBody, RoleQueryData, RoleUpdateBody } from '~/types/role'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useRoleApi() {
  async function queryRoles(body: RoleQueryBody) {
    const response = await $fetch<ApiResponse<RoleQueryData>>('/api/role/query', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取角色列表失败')
  }

  async function listRoles(body: RoleListBody = {}) {
    const response = await $fetch<ApiResponse<RoleOption[]>>('/api/role/list', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '获取角色列表失败')
  }

  async function addRole(body: RoleAddBody) {
    const response = await $fetch<ApiResponse<Role>>('/api/role/add', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '新增角色失败')
  }

  async function updateRole(body: RoleUpdateBody) {
    const response = await $fetch<ApiResponse<Role>>('/api/role/update', {
      method: 'POST',
      body
    })
    return ensureResponseSuccess(response, '更新角色失败')
  }

  async function deleteRole(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/role/delete/${id}`, {
      method: 'POST'
    })
    return ensureResponseSuccess(response, '删除角色失败')
  }

  return {
    queryRoles,
    listRoles,
    addRole,
    updateRole,
    deleteRole
  }
}
