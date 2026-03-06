import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Transition, TransitionAddBody } from '~/types/transition'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TransitionAddBody>(event)

  const response = await backendFetch<ApiResponse<Transition>>(event, '/api/transition/add', {
    method: 'POST',
    body: {
      process_id: body?.process_id,
      from_node_id: body?.from_node_id,
      to_node_id: body?.to_node_id
    }
  })

  return ensureApiSuccess(response, 'Transition add failed', 502)
})
