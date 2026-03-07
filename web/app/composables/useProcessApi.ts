import type { ApiResponse } from '~/types/api'
import type { Process, ProcessAddBody, ProcessGetData, ProcessQueryBody, ProcessQueryData, ProcessUpdateBody } from '~/types/process'

function ensureResponseSuccess<T>(response: ApiResponse<T>, fallbackMessage: string): T {
  if (response.code !== 0) {
    throw new Error(response.message || fallbackMessage)
  }

  return response.data
}

export function useProcessApi() {
  async function getProcess(id: number) {
    const response = await $fetch<ApiResponse<ProcessGetData>>(`/api/process/get/${id}`, {
      method: 'GET'
    })

    return ensureResponseSuccess(response, '获取流程详情失败')
  }

  async function queryProcess(body: ProcessQueryBody) {
    const response = await $fetch<ApiResponse<ProcessQueryData>>('/api/process/query', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '查询流程失败')
  }

  async function addProcess(body: ProcessAddBody) {
    const response = await $fetch<ApiResponse<Process>>('/api/process/add', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '创建流程失败')
  }

  async function updateProcess(body: ProcessUpdateBody) {
    const response = await $fetch<ApiResponse<Process>>('/api/process/update', {
      method: 'POST',
      body
    })

    return ensureResponseSuccess(response, '更新流程失败')
  }

  async function deleteProcess(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/process/delete/${id}`, {
      method: 'POST'
    })

    return ensureResponseSuccess(response, '删除流程失败')
  }

  async function disableProcess(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/process/disable/${id}`, {
      method: 'POST'
    })

    return ensureResponseSuccess(response, '禁用流程失败')
  }

  async function enableProcess(id: number) {
    const response = await $fetch<ApiResponse<null>>(`/api/process/enable/${id}`, {
      method: 'POST'
    })

    return ensureResponseSuccess(response, '启用流程失败')
  }

  return {
    getProcess,
    queryProcess,
    addProcess,
    updateProcess,
    deleteProcess,
    disableProcess,
    enableProcess
  }
}
