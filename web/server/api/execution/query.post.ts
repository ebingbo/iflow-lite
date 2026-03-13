import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { ExecutionQueryBody, ExecutionQueryData } from '~/types/execution'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<ExecutionQueryBody>(event)

  const response = await backendFetch<ApiResponse<ExecutionQueryData>>(event, '/api/execution/query', {
    method: 'POST',
    body: {
      page: body?.page,
      size: body?.size,
      process_id: body?.process_id,
      process_code: body?.process_code,
      business_key: body?.business_key,
      business_type: body?.business_type,
      status: body?.status
    }
  })

  return ensureApiSuccess(response, 'Execution query failed', 502)
})
