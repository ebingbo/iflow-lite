import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Node, NodeUpdateBody } from '~/types/node'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<NodeUpdateBody>(event)

  const response = await backendFetch<ApiResponse<Node>>(event, '/api/node/update', {
    method: 'POST',
    body: {
      id: body?.id,
      tag: body?.tag,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Node update failed', 502)
})
