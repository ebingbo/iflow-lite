import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskQueryBody, TaskQueryData } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskQueryBody>(event)

  const response = await backendFetch<ApiResponse<TaskQueryData>>(event, '/api/task/query', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      process_id: body?.process_id,
      process_code: body?.process_code,
      execution_id: body?.execution_id,
      node_id: body?.node_id,
      node_code: body?.node_code,
      assignee_id: body?.assignee_id,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'Task query failed', 502)
})
