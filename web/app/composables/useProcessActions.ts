import type { ProcessAddBody, ProcessUpdateBody } from '~/types/process'

interface ProcessActionFeedbackOptions {
  refreshKey?: string
  successTitle?: string
  successDescription?: string
  errorTitle?: string
}

type ProcessToggleAction = 'enable' | 'disable'

export function useProcessActions() {
  const toast = useToast()
  const {
    addProcess,
    updateProcess,
    deleteProcess,
    disableProcess,
    enableProcess
  } = useProcessApi()

  async function addProcessWithFeedback(body: ProcessAddBody, options: ProcessActionFeedbackOptions = {}) {
    try {
      const result = await addProcess(body)
      toast.add({
        title: options.successTitle || '成功',
        description: options.successDescription || `新的流程 ${result.name} 已添加`,
        color: 'success',
        icon: 'i-lucide-circle-check'
      })
      await refreshNuxtData(options.refreshKey || 'process-query')
      return result
    } catch (error) {
      toast.add({
        title: options.errorTitle || '创建失败',
        description: error instanceof Error ? error.message : '请求失败，请稍后重试',
        color: 'error',
        icon: 'i-lucide-circle-x'
      })
      throw error
    }
  }

  async function updateProcessWithFeedback(body: ProcessUpdateBody, options: ProcessActionFeedbackOptions = {}) {
    try {
      const result = await updateProcess(body)
      toast.add({
        title: options.successTitle || '更新成功',
        description: options.successDescription || `流程 ${result.name} 已更新`,
        color: 'success',
        icon: 'i-lucide-circle-check'
      })
      await refreshNuxtData(options.refreshKey || 'process-query')
      return result
    } catch (error) {
      toast.add({
        title: options.errorTitle || '更新失败',
        description: error instanceof Error ? error.message : '请求失败，请稍后重试',
        color: 'error',
        icon: 'i-lucide-circle-x'
      })
      throw error
    }
  }

  async function deleteProcessWithFeedback(id: number, options: ProcessActionFeedbackOptions = {}) {
    try {
      await deleteProcess(id)
      toast.add({
        title: options.successTitle || '删除成功',
        description: options.successDescription,
        color: 'success',
        icon: 'i-lucide-circle-check'
      })
      await refreshNuxtData(options.refreshKey || 'process-query')
    } catch (error) {
      toast.add({
        title: options.errorTitle || '删除失败',
        description: error instanceof Error ? error.message : '请求失败，请稍后重试',
        color: 'error',
        icon: 'i-lucide-circle-x'
      })
      throw error
    }
  }

  async function toggleProcessStatusWithFeedback(
    id: number,
    action: ProcessToggleAction,
    options: ProcessActionFeedbackOptions = {}
  ) {
    const actionText = action === 'enable' ? '启用' : '禁用'

    try {
      if (action === 'enable') {
        await enableProcess(id)
      } else {
        await disableProcess(id)
      }

      toast.add({
        title: options.successTitle || `${actionText}成功`,
        description: options.successDescription,
        color: 'success',
        icon: 'i-lucide-circle-check'
      })
      await refreshNuxtData(options.refreshKey || 'process-query')
    } catch (error) {
      toast.add({
        title: options.errorTitle || `${actionText}失败`,
        description: error instanceof Error ? error.message : '请求失败，请稍后重试',
        color: 'error',
        icon: 'i-lucide-circle-x'
      })
      throw error
    }
  }

  return {
    addProcessWithFeedback,
    updateProcessWithFeedback,
    deleteProcessWithFeedback,
    toggleProcessStatusWithFeedback
  }
}
