import type { ApiResponse } from '~/types/api'
import type { Transition, TransitionAddBody, TransitionListBody } from '~/types/transition'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useTransitionApi() {
  async function addTransition(body: TransitionAddBody) {
    const response = await $fetch<ApiResponse<Transition>>('/api/transition/add', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '新增连线失败')
  }

  async function listTransition(body: TransitionListBody) {
    const response = await $fetch<ApiResponse<Transition[]>>('/api/transition/list', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '获取连线列表失败')
  }

  return {
    addTransition,
    listTransition
  }
}
