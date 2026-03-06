import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Transition, TransitionListBody } from '~/types/transition'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TransitionListBody>(event)

  const response = await backendFetch<ApiResponse<Transition[]>>(event, '/api/transition/list', {
    method: 'POST',
    body: {
      process_id: body?.process_id
    }
  })

  return ensureApiSuccess(response, 'Transition list failed', 502)
})
