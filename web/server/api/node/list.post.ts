import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Node, NodeListBody } from '~/types/node'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<NodeListBody>(event)

  const response = await backendFetch<ApiResponse<Node[]>>(event, '/api/node/list', {
    method: 'POST',
    body: {
      process_id: body?.process_id,
      process_code: body?.process_code,
      code: body?.code,
      type: body?.type
    }
  })

  return ensureApiSuccess(response, 'Node list failed', 502)
})
