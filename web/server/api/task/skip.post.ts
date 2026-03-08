import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskSkipBody } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskSkipBody>(event)

  const response = await backendFetch<ApiResponse<null>>(event, '/api/task/skip', {
    method: 'POST',
    body: {
      id: body?.id,
      assignee_id: body?.assignee_id
    }
  })

  return ensureApiSuccess(response, 'Task skip failed', 502)
})
