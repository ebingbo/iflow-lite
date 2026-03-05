import { readBody } from 'h3'
import type { ApiMessageResponse } from '~/types/api'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'
import type { ProcessDeleteBody } from '~/types/process'

export default eventHandler(async (event) => {
  const body = await readBody<ProcessDeleteBody>(event)
  const ids = Array.isArray(body?.ids) ? body.ids.filter(id => Number.isInteger(id)) : []

  if (!ids.length) {
    throw createError({
      statusCode: 400,
      statusMessage: 'ids is required'
    })
  }

  const response = await backendFetch<ApiMessageResponse>(event, '/api/process/delete', {
    method: 'POST',
    body: {
      ids
    }
  })

  return ensureApiSuccess(response, 'Process delete failed', 502)
})
