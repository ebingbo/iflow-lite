import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Transition, TransitionUpdateBody } from '~/types/transition'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TransitionUpdateBody>(event)

  const response = await backendFetch<ApiResponse<Transition>>(event, '/api/transition/update', {
    method: 'POST',
    body: {
      id: body?.id,
      from_node_id: body?.from_node_id,
      to_node_id: body?.to_node_id
    }
  })

  return ensureApiSuccess(response, 'Transition update failed', 502)
})
