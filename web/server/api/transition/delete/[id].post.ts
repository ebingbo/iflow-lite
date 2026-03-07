import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid transition id'
    })
  }

  const response = await backendFetch<ApiResponse<null>>(event, `/api/transition/delete/${id}`, {
    method: 'POST'
  })

  return ensureApiSuccess(response, 'Transition delete failed', 502)
})
