import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskClaimableQueryBody, TaskQueryData } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskClaimableQueryBody>(event)

  const response = await backendFetch<ApiResponse<TaskQueryData>>(event, '/api/task/query/claimable', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'Task claimable query failed', 502)
})
