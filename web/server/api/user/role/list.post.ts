import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserRoleListBody, UserRoleListData } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserRoleListBody>(event)

  const response = await backendFetch<ApiResponse<UserRoleListData>>(event, '/api/user/role/list', {
    method: 'POST',
    body: {
      user_id: body?.user_id
    }
  })

  return ensureApiSuccess(response, 'User role list failed', 502)
})
