import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserStatusUpdateBody, User } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserStatusUpdateBody>(event)

  const response = await backendFetch<ApiResponse<User>>(event, '/api/user/status/update', {
    method: 'POST',
    body: {
      id: body?.id,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'User status update failed', 502)
})
