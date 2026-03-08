import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserListBody, UserOption } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserListBody>(event)

  const response = await backendFetch<ApiResponse<UserOption[]>>(event, '/api/user/list', {
    method: 'POST',
    body: {
      keyword: body?.keyword,
      size: body?.size
    }
  })

  return ensureApiSuccess(response, 'User list failed', 502)
})
