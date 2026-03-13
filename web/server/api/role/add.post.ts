import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Role, RoleAddBody } from '~/types/role'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<RoleAddBody>(event)

  const response = await backendFetch<ApiResponse<Role>>(event, '/api/role/add', {
    method: 'POST',
    body: {
      name: body?.name,
      code: body?.code
    }
  })

  return ensureApiSuccess(response, 'Role add failed', 502)
})
