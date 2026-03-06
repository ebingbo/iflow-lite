import type { ProcessQueryData } from '~/types/process'

interface ProcessListOptions {
  pageSize?: number
}

export async function useProcessList(options: ProcessListOptions = {}) {
  const { queryProcess } = useProcessApi()
  const pagination = ref({
    pageIndex: 0,
    pageSize: options.pageSize || 10
  })

  const requestBody = computed(() => ({
    page: pagination.value.pageIndex + 1,
    size: pagination.value.pageSize
  }))

  const { data, status } = await useAsyncData<ProcessQueryData>(
    'process-query',
    () => queryProcess(requestBody.value),
    {
      watch: [requestBody],
      default: () => ({
        total: 0,
        items: []
      }),
      lazy: true
    }
  )

  const items = computed(() => data.value?.items || [])
  const total = computed(() => data.value?.total || 0)
  const loading = computed(() => ['idle', 'pending'].includes(status.value))

  async function refresh() {
    await refreshNuxtData('process-query')
  }

  return {
    pagination,
    requestBody,
    items,
    total,
    loading,
    refresh
  }
}
