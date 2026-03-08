import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskLogQueryBody, TaskLogQueryData } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskLogQueryBody>(event)

  const response = await backendFetch<ApiResponse<TaskLogQueryData>>(event, '/api/log/query', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      task_id: body?.task_id
    }
  })

  return ensureApiSuccess(response, 'Task logs query failed', 502)
})
