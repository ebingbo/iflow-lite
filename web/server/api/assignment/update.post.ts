import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Assignment, AssignmentUpdateBody } from '~/types/assignment'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<AssignmentUpdateBody>(event)

  const response = await backendFetch<ApiResponse<Assignment>>(event, '/api/assignment/update', {
    method: 'POST',
    body: {
      id: body?.id,
      type: body?.type,
      value: body?.value,
      priority: body?.priority,
      strategy: body?.strategy
    }
  })

  return ensureApiSuccess(response, 'Assignment update failed', 502)
})
