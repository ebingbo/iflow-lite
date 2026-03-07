import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid node id'
    })
  }

  const response = await backendFetch<ApiResponse<null>>(event, `/api/node/delete/${id}`, {
    method: 'POST'
  })

  return ensureApiSuccess(response, 'Node delete failed', 502)
})
