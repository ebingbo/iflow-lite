import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserPasswordForgotBody } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserPasswordForgotBody>(event)

  const response = await backendFetch<ApiResponse<null>>(event, '/api/user/password/forgot', {
    method: 'POST',
    body: {
      email: body?.email
    }
  })

  return ensureApiSuccess(response, 'Forgot password failed', 400)
})
