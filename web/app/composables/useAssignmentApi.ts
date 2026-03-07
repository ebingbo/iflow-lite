import type { ApiResponse } from '~/types/api'
import type { Assignment, AssignmentAddBody, AssignmentListBody, AssignmentUpdateBody } from '~/types/assignment'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useAssignmentApi() {
  async function addAssignment(body: AssignmentAddBody) {
    const response = await $fetch<ApiResponse<Assignment>>('/api/assignment/add', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '新增分派失败')
  }

  async function updateAssignment(body: AssignmentUpdateBody) {
    const response = await $fetch<ApiResponse<Assignment>>('/api/assignment/update', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '更新分派失败')
  }

  async function listAssignment(body: AssignmentListBody) {
    const response = await $fetch<ApiResponse<Assignment[]>>('/api/assignment/list', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '获取分派列表失败')
  }

  async function deleteAssignment(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/assignment/delete/${id}`, {
      method: 'POST'
    })

    return ensureResponseSuccess(response, '删除分派失败')
  }

  return {
    addAssignment,
    updateAssignment,
    listAssignment,
    deleteAssignment
  }
}
