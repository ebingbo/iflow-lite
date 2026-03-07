import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Assignment, AssignmentListBody } from '~/types/assignment'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<AssignmentListBody>(event)

  const response = await backendFetch<ApiResponse<Assignment[]>>(event, '/api/assignment/list', {
    method: 'POST',
    body: {
      process_id: body?.process_id,
      node_id: body?.node_id
    }
  })

  return ensureApiSuccess(response, 'Assignment list failed', 502)
})
