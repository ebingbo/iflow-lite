import type { ApiResponse } from '~/types/api'
import type { Node, NodeAddBody, NodeListBody } from '~/types/node'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useNodeApi() {
  async function addNode(body: NodeAddBody) {
    const response = await $fetch<ApiResponse<Node>>('/api/node/add', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '新增节点失败')
  }

  async function listNode(body: NodeListBody) {
    const response = await $fetch<ApiResponse<Node[]>>('/api/node/list', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '获取节点列表失败')
  }

  return {
    addNode,
    listNode
  }
}
