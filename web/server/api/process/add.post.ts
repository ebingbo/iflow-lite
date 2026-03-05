import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'
import type { Process, ProcessAddBody } from '~/types/process'

export default eventHandler(async (event) => {
  const body = await readBody<ProcessAddBody>(event)

  const response = await backendFetch<ApiResponse<Process>>(event, '/api/process/add', {
    method: 'POST',
    body: {
      code: body?.code,
      name: body?.name,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Process add failed', 502)
})
