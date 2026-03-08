import type { ApiResponse } from '~/types/api'
import type {
  TaskCandidateListBody,
  TaskCandidateListData,
  TaskClaimBody,
  TaskClaimableQueryBody,
  TaskCompleteBody,
  TaskLogQueryBody,
  TaskLogQueryData,
  TaskQueryBody,
  TaskQueryData,
  TaskSkipBody
} from '~/types/task'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useTaskApi() {
  async function queryTasks(body: TaskQueryBody) {
    const response = await $fetch<ApiResponse<TaskQueryData>>('/api/task/query', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询任务失败')
  }

  async function queryClaimableTasks(body: TaskClaimableQueryBody) {
    const response = await $fetch<ApiResponse<TaskQueryData>>('/api/task/query/claimable', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询可认领任务失败')
  }

  async function claimTask(body: TaskClaimBody) {
    const response = await $fetch<ApiResponse<null>>('/api/task/claim', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '认领任务失败')
  }

  async function completeTask(body: TaskCompleteBody) {
    const response = await $fetch<ApiResponse<null>>('/api/task/complete', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '完成任务失败')
  }

  async function skipTask(body: TaskSkipBody) {
    const response = await $fetch<ApiResponse<null>>('/api/task/skip', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '跳过任务失败')
  }

  async function listTaskCandidates(body: TaskCandidateListBody) {
    const response = await $fetch<ApiResponse<TaskCandidateListData>>('/api/task/candidate/list', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询任务候选人失败')
  }

  async function queryTaskLogs(body: TaskLogQueryBody) {
    const response = await $fetch<ApiResponse<TaskLogQueryData>>('/api/log/query', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询任务日志失败')
  }

  return {
    queryTasks,
    queryClaimableTasks,
    claimTask,
    completeTask,
    skipTask,
    listTaskCandidates,
    queryTaskLogs
  }
}
