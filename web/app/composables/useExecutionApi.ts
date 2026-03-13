import type { ApiResponse } from '~/types/api'
import type { Execution, ExecutionQueryBody, ExecutionQueryData, ExecutionStartBody } from '~/types/execution'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useExecutionApi() {
  async function queryExecutions(body: ExecutionQueryBody) {
    const response = await $fetch<ApiResponse<ExecutionQueryData>>('/api/execution/query', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询实例失败')
  }

  async function getExecution(id: number) {
    const response = await $fetch<ApiResponse<Execution>>(`/api/execution/get/${id}`, {
      method: 'GET'
    })

    return ensureResponseSuccess(response, '查询实例详情失败')
  }

  async function startExecution(body: ExecutionStartBody) {
    const response = await $fetch<ApiResponse<Execution>>('/api/execution/start', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '发起流程失败')
  }

  return {
    queryExecutions,
    getExecution,
    startExecution
  }
}
