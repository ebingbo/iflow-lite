import type { ApiResponse } from '~/types/api'
import type { Execution } from '~/types/execution'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = event.context.params?.id
  const response = await backendFetch<ApiResponse<Execution>>(event, `/api/execution/get?id=${id}`, {
    method: 'GET'
  })

  return ensureApiSuccess(response, 'Execution get failed', 502)
})
