import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid process id'
    })
  }

  const response = await backendFetch<ApiResponse<null>>(event, `/api/process/enable/${id}`, {
    method: 'POST'
  })

  return ensureApiSuccess(response, 'Process enable failed', 502)
})
