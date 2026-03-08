<script setup lang="ts">
useSeoMeta({
  title: '仪表盘',
  description: '流程系统概览'
})
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="仪表盘"
      description="查看待办、流程运行与系统概览"
    />

    <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
      <UCard>
        <p class="text-sm text-muted">
          我的待办
        </p>
        <p class="mt-2 text-2xl font-semibold">
          -
        </p>
      </UCard>
      <UCard>
        <p class="text-sm text-muted">
          我发起的流程
        </p>
        <p class="mt-2 text-2xl font-semibold">
          -
        </p>
      </UCard>
      <UCard>
        <p class="text-sm text-muted">
          运行中实例
        </p>
        <p class="mt-2 text-2xl font-semibold">
          -
        </p>
      </UCard>
      <UCard>
        <p class="text-sm text-muted">
          逾期任务
        </p>
        <p class="mt-2 text-2xl font-semibold">
          -
        </p>
      </UCard>
    </div>
  </UContainer>
</template>
