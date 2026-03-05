import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserProfile, UserSignupBody } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserSignupBody>(event)

  const response = await backendFetch<ApiResponse<UserProfile>>(event, '/api/user/add', {
    method: 'POST',
    body: {
      name: body?.name,
      email: body?.email,
      password: body?.password
    }
  })

  return ensureApiSuccess(response, 'Signup failed', 401)
})
