import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserRoleUpdateBody } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserRoleUpdateBody>(event)

  const response = await backendFetch<ApiResponse<null>>(event, '/api/user/role/update', {
    method: 'POST',
    body: {
      user_id: body?.user_id,
      role_ids: body?.role_ids || []
    }
  })

  return ensureApiSuccess(response, 'User role update failed', 502)
})
