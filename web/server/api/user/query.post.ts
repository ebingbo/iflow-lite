import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserQueryBody, UserQueryData } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserQueryBody>(event)

  const response = await backendFetch<ApiResponse<UserQueryData>>(event, '/api/user/query', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      email: body?.email,
      name: body?.name,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'User query failed', 502)
})
