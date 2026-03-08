import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskCompleteBody } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskCompleteBody>(event)

  const response = await backendFetch<ApiResponse<null>>(event, '/api/task/complete', {
    method: 'POST',
    body: {
      id: body?.id,
      assignee_id: body?.assignee_id,
      remark: body?.remark || ''
    }
  })

  return ensureApiSuccess(response, 'Task complete failed', 502)
})
