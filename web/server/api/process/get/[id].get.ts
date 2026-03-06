import type { ApiResponse } from '~/types/api'
import type { Process } from '~/types/process'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = getRouterParam(event, 'id')

  if (!id || !/^\d+$/.test(id)) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid process id'
    })
  }

  const response = await backendFetch<ApiResponse<Process>>(event, `/api/process/get?id=${id}`, {
    method: 'GET'
  })

  return ensureApiSuccess(response, 'Process get failed', 502)
})
