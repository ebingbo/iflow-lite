import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskClaimBody } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskClaimBody>(event)

  const response = await backendFetch<ApiResponse<null>>(event, '/api/task/claim', {
    method: 'POST',
    body: {
      id: body?.id
    }
  })

  return ensureApiSuccess(response, 'Task claim failed', 502)
})
