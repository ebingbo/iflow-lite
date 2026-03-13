import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { RoleQueryBody, RoleQueryData } from '~/types/role'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<RoleQueryBody>(event)

  const response = await backendFetch<ApiResponse<RoleQueryData>>(event, '/api/role/query', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      keyword: body?.keyword
    }
  })

  return ensureApiSuccess(response, 'Role query failed', 502)
})
