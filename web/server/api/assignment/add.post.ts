import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Assignment, AssignmentAddBody } from '~/types/assignment'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<AssignmentAddBody>(event)

  const response = await backendFetch<ApiResponse<Assignment>>(event, '/api/assignment/add', {
    method: 'POST',
    body: {
      process_id: body?.process_id,
      node_id: body?.node_id,
      type: body?.type,
      value: body?.value,
      priority: body?.priority,
      strategy: body?.strategy
    }
  })

  return ensureApiSuccess(response, 'Assignment add failed', 502)
})
