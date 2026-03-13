import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { Execution, ExecutionStartBody } from '~/types/execution'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<ExecutionStartBody>(event)

  const response = await backendFetch<ApiResponse<Execution>>(event, '/api/execution/start', {
    method: 'POST',
    body: {
      process_code: body?.process_code,
      business_key: body?.business_key,
      business_type: body?.business_type,
      created_by: body?.created_by
    }
  })

  return ensureApiSuccess(response, 'Execution start failed', 502)
})
