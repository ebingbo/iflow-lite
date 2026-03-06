import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Node, NodeAddBody } from '~/types/node'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<NodeAddBody>(event)

  const response = await backendFetch<ApiResponse<Node>>(event, '/api/node/add', {
    method: 'POST',
    body: {
      process_id: body?.process_id,
      process_code: body?.process_code,
      tag: body?.tag,
      name: body?.name,
      code: body?.code,
      type: body?.type,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Node add failed', 502)
})
