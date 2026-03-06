import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'
import type { Process, ProcessUpdateBody } from '~/types/process'

export default eventHandler(async (event) => {
  const body = await readBody<ProcessUpdateBody>(event)

  if (!Number.isInteger(body?.id) || body.id <= 0) {
    throw createError({
      statusCode: 400,
      statusMessage: 'invalid process id'
    })
  }

  const response = await backendFetch<ApiResponse<Process>>(event, '/api/process/update', {
    method: 'POST',
    body: {
      id: body.id,
      name: body?.name,
      description: body?.description
    }
  })

  return ensureApiSuccess(response, 'Process update failed', 502)
})
