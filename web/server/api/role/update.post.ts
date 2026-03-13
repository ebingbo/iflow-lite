import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Role, RoleUpdateBody } from '~/types/role'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<RoleUpdateBody>(event)

  const response = await backendFetch<ApiResponse<Role>>(event, '/api/role/update', {
    method: 'POST',
    body: {
      id: body?.id,
      name: body?.name,
      code: body?.code
    }
  })

  return ensureApiSuccess(response, 'Role update failed', 502)
})
