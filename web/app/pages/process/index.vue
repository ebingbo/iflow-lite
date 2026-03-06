<script setup lang="ts">
import type { TableColumn } from '#ui/components/Table.vue'
import type { Process } from '~/types/process'
import dayjs from 'dayjs'
import type { Row } from '@tanstack/vue-table'
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
const UDropdownMenu = resolveComponent('UDropdownMenu')
const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')
const { pagination, items: tableData, total, loading, refresh } = await useProcessList()
const { deleteProcessWithFeedback, toggleProcessStatusWithFeedback } = useProcessActions()
const editOpen = ref(false)
const editingProcess = ref<Process | null>(null)

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
      label: '查看',
      onSelect() {
        navigateTo(`/process/${row.original.id}`)
      }
    },
    {
      label: '编辑',
      onSelect() {
        openEditModal(row.original)
      }
    },
    {
      type: 'separator'
    }

  ]
  if (row.original.status === 1) {
    items.push({
      label: '禁用',
      async onSelect() {
        await processAction(row.original.id, 'disable')
      }
    })
  } else {
    items.push({
      label: '启用',
      async onSelect() {
        await processAction(row.original.id, 'enable')
      }
    })
    items.push({
      label: '删除',
      async onSelect() {
        await processAction(row.original.id, 'delete')
      }
    })
  }
  return items
}

const globalFilter = ref('')

function openEditModal(process: Process) {
  editingProcess.value = process
  editOpen.value = true
}

async function refreshProcessQuery() {
  await refresh()
}

async function processAction(id: number, action: 'enable' | 'disable' | 'delete') {
  try {
    if (action === 'delete') {
      await deleteProcessWithFeedback(id)
      return
    }

    await toggleProcessStatusWithFeedback(id, action)
  } catch {
    return
  }
}

async function handleEditSuccess() {
  editingProcess.value = null
  await refreshProcessQuery()
}

function handleEditOpenUpdate(value: boolean) {
  editOpen.value = value
  if (!value) {
    editingProcess.value = null
  }
}
</script>

<template>
  <UContainer>
    <UPageHeader
      title="流程定义管理"
      description="集中管理流程定义，支持查看、筛选与维护流程状态。"
      :links="links"
    />
    <div class="flex justify-between py-3.5 border-b border-accented">
      <UInput
        v-model="globalFilter"
        class="max-w-sm"
        placeholder="Filter..."
      />
      <ProcessAddModal />
    </div>
    <UTable
      v-model:global-filter="globalFilter"
      sticky
      :loading="loading"
      loading-color="primary"
      loading-animation="carousel"
      :data="tableData"
      :columns="columns"
      class="flex-1"
    />
    <div class="flex justify-end border-t border-default pt-4 px-4">
      <UPagination
        :page="(pagination.pageIndex || 0) + 1"
        :items-per-page="pagination.pageSize"
        :total="total"
        @update:page="(p: number) => { pagination.pageIndex = p - 1 }"
      />
    </div>
    <ProcessEditModal
      v-model:open="editOpen"
      :process="editingProcess"
      @success="handleEditSuccess"
      @update:open="handleEditOpenUpdate"
    />
  </UContainer>
</template>
