<script setup lang="ts">
const columns = [{
  label: '产品',
  children: [{
    label: '流程管理',
    to: '/process'
  }, {
    label: '文档中心',
    to: '/docs'
  }, {
    label: '更新日志',
    to: '/changelog'
  }, {
    label: '博客',
    to: '/blog'
  }]
}, {
  label: '能力',
  children: [{
    label: '流程定义'
  }, {
    label: '任务执行'
  }, {
    label: '权限角色'
  }, {
    label: '日志审计'
  }]
}, {
  label: '项目',
  children: [{
    label: '代码仓库',
    to: 'https://github.com/ebingbo/iflow-lite',
    target: '_blank'
  }, {
    label: '开源协议',
    to: '/docs',
    target: '_blank'
  }, {
    label: '问题反馈',
    to: 'https://github.com/ebingbo/iflow-lite/issues',
    target: '_blank'
  }, {
    label: '参与贡献',
    to: 'https://github.com/ebingbo/iflow-lite',
    target: '_blank'
  }]
}]
</script>

<template>
  <USeparator
    icon="i-lucide-workflow"
    class="h-px"
  />

  <UFooter :ui="{ top: 'border-b border-default' }">
    <template #top>
      <UContainer>
        <UFooterColumns :columns="columns">
          <template #right>
            <div class="space-y-4 max-w-xs">
              <p class="text-sm text-muted">
                iFlow Lite 是一个轻量级流程管理与工作流引擎，适合快速构建企业内部流程系统。
              </p>
              <div class="flex gap-2">
                <UButton
                  to="/process"
                  icon="i-lucide-arrow-right"
                  label="进入流程管理"
                  size="sm"
                />
                <UButton
                  to="/docs"
                  label="查看文档"
                  color="neutral"
                  variant="subtle"
                  size="sm"
                />
              </div>
            </div>
          </template>
        </UFooterColumns>
      </UContainer>
    </template>

    <template #left>
      <p class="text-muted text-sm">
        iFlow Lite • Apache-2.0 • © {{ new Date().getFullYear() }}
      </p>
    </template>

    <template #right>
      <UButton
        to="https://github.com/ebingbo/iflow-lite/issues"
        target="_blank"
        icon="i-lucide-message-circle-warning"
        aria-label="iFlow Lite Issues"
        color="neutral"
        variant="ghost"
      />
      <UButton
        to="/docs"
        icon="i-lucide-book-open-text"
        aria-label="iFlow Lite Docs"
        color="neutral"
        variant="ghost"
      />
      <UButton
        to="https://github.com/ebingbo/iflow-lite"
        target="_blank"
        icon="i-simple-icons-github"
        aria-label="iFlow Lite on GitHub"
        color="neutral"
        variant="ghost"
      />
    </template>
  </UFooter>
</template>
