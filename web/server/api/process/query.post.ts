import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'
import type { ProcessQueryBody, ProcessQueryData } from '~/types/process'

export default eventHandler(async (event) => {
  const body = await readBody<ProcessQueryBody>(event)

  const page = Number(body?.page) > 0 ? Number(body?.page) : 1
  const size = Number(body?.size) > 0 ? Number(body?.size) : 10

  const response = await backendFetch<ApiResponse<ProcessQueryData>>(event, '/api/process/query', {
    method: 'POST',
    body: {
      page,
      size,
      name: body?.name,
      code: body?.code,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'Process query failed', 502)
})
