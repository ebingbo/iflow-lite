import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { RoleListBody, RoleOption } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<RoleListBody>(event)

  const response = await backendFetch<ApiResponse<RoleOption[]>>(event, '/api/role/list', {
    method: 'POST',
    body: {
      keyword: body?.keyword,
      size: body?.size
    }
  })

  return ensureApiSuccess(response, 'Role list failed', 502)
})
