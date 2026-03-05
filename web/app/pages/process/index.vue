<script setup lang="ts">
import type { TableColumn } from '#ui/components/Table.vue'
import type { Process, ProcessQueryData } from '~/types/process'
import type { ApiResponse } from '~/types/api'
import dayjs from 'dayjs'
import type { Row, Table } from '@tanstack/vue-table'
import { getPaginationRowModel } from '@tanstack/vue-table'
import type { ButtonProps } from '@nuxt/ui'
import type { DropdownMenuItem } from '#ui/components/DropdownMenu.vue'

const title = '流程'
const description = '流程定义'

useSeoMeta({
  title,
  ogTitle: title,
  description,
  ogDescription: description
})

const links = ref<ButtonProps[]>([
  {
    variant: 'ghost',
    icon: 'i-lucide-circle-question-mark',
    to: 'https://github.com/nuxt/ui/tree/v4/src/runtime/components/PageHeader.vue',
    target: '_blank'
  }
])
const toast = useToast()
const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')
const table = useTemplateRef<{ tableApi?: Table<Process> }>('table')
const pagination = ref({
  pageIndex: 0,
  pageSize: 10
})

const requestBody = computed(() => ({
  page: pagination.value.pageIndex + 1,
  size: pagination.value.pageSize
}))

const { data } = await useFetch<ApiResponse<ProcessQueryData>>('/api/process/query', {
  key: 'process-query',
  method: 'POST',
  body: requestBody,
  watch: [requestBody],
  default: () => ({
    code: 0,
    message: 'ok',
    data: {
      total: 0,
      items: []
    }
  }),
  lazy: true
})

const tableData = computed(() => data.value?.data?.items || [])
const total = computed(() => data.value?.data?.total || 0)

const columns: TableColumn<Process>[] = [
  {
    accessorKey: 'id',
    header: '#',
    cell: ({ row }) => `#${row.getValue('id')}`
  },
  {
    accessorKey: 'name',
    header: '名称'
  },
  {
    accessorKey: 'code',
    header: '编码'
  },
  {
    accessorKey: 'description',
    header: '描述'
  },
  {
    accessorKey: 'status',
    header: '状态',
    cell: ({ row }) => {
      const color = {
        1: 'success' as const,
        0: 'neutral' as const
      }[row.getValue('status') as number]

      return h(UBadge, { class: 'capitalize', variant: 'subtle', color }, () =>
        row.getValue('status') === 1 ? '正常' : '禁用'
      )
    }
  },
  {
    accessorKey: 'created_by_name',
    header: '创建用户'
  },
  {
    accessorKey: 'updated_by_name',
    header: '更新用户'
  },
  {
    accessorKey: 'created_at',
    header: '创建时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('created_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    accessorKey: 'updated_at',
    header: '更新时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('updated_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    id: 'actions',
    header: '操作',
    meta: {
      class: {
        td: 'text-right'
      }
    },
    cell: ({ row }) => {
      return h(
        UDropdownMenu,
        {
          items: getRowItems(row)
        },
        () =>
          h(UButton, {
            'icon': 'i-lucide-ellipsis-vertical',
            'color': 'neutral',
            'variant': 'ghost',
            'aria-label': 'Actions dropdown'
          })
      )
    }
  }
]

function getRowItems(row: Row<Process>) {
  const items: DropdownMenuItem[] = [
    {
      label: '查看'
    },
    {
      label: '编辑'
    },
    {
      type: 'separator'
    }
  ]
  if (row.original.status === 1) {
    items.push({
      label: '禁用',
      onSelect() {
        toast.add({
          title: '已禁用',
          color: 'success',
          icon: 'i-lucide-circle-check'
        })
      }
    })
  } else {
    items.push({
      label: '启用'
    })
  }
  return items
}

const globalFilter = ref('')
</script>

<template>
  <UContainer>
    <UPageHeader
      title="工作流程"
      description="A responsive page header with title, description and actions."
      :links="links"
    />
    <div class="flex justify-between py-3.5 border-b border-accented">
      <UInput
        v-model="globalFilter"
        class="max-w-sm"
        placeholder="Filter..."
      />
      <UButton
        icon="i-lucide-plus"
        size="md"
        color="primary"
        variant="solid"
      >
        新建流程
      </UButton>
    </div>
    <UTable
      ref="table"
      v-model:global-filter="globalFilter"
      sticky
      loading
      loading-color="primary"
      loading-animation="carousel"
      :data="tableData"
      :columns="columns"
      :pagination-options="{
        getPaginationRowModel: getPaginationRowModel()
      }"
      class="flex-1"
    />
    <div class="flex justify-end border-t border-default pt-4 px-4">
      <UPagination
        :page="(table?.tableApi?.getState().pagination.pageIndex || 0) + 1"
        :items-per-page="table?.tableApi?.getState().pagination.pageSize"
        :total="total"
        @update:page="(p: number) => table?.tableApi?.setPageIndex(p - 1)"
      />
    </div>
  </UContainer>
</template>
