import type { ApiResponse } from '~/types/api'
import type { Node, NodeAddBody, NodeListBody, NodeUpdateBody } from '~/types/node'

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

  async function updateNode(body: NodeUpdateBody) {
    const response = await $fetch<ApiResponse<Node>>('/api/node/update', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '更新节点失败')
  }

  async function deleteNode(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/node/delete/${id}`, {
      method: 'POST'
    })

    return ensureResponseSuccess(response, '删除节点失败')
  }

  return {
    addNode,
    listNode,
    updateNode,
    deleteNode
  }
}
