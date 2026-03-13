import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const id = event.context.params?.id
  const response = await backendFetch<ApiResponse<null>>(event, `/api/role/delete/${id}`, {
    method: 'POST'
  })

  return ensureApiSuccess(response, 'Role delete failed', 502)
})
