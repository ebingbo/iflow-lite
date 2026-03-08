import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { UserProfile, UserProfileUpdateBody } from '~/types/user'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<UserProfileUpdateBody>(event)

  const response = await backendFetch<ApiResponse<UserProfile>>(event, '/api/user/profile/update', {
    method: 'POST',
    body: {
      name: body?.name
    }
  })

  return ensureApiSuccess(response, 'Update profile failed', 400)
})
