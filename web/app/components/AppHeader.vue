<script setup lang="ts">
const route = useRoute()
const { user, logout } = useAuth()

const publicItems = computed(() => [
  {
    label: '产品介绍',
    to: '/'
  },
  {
    label: '文档中心',
    to: '/docs',
    active: route.path.startsWith('/docs')
  }, {
    label: '价格',
    to: '/pricing'
  }, {
    label: '博客',
    to: '/blog'
  }, {
    label: '更新日志',
    to: '/changelog'
  }])

const consoleItems = computed(() => [
  {
    label: '仪表盘',
    to: '/dashboard'
  },
  {
    label: '流程管理',
    to: '/process'
  },
  {
    label: '任务中心',
    to: '/tasks'
  },
  {
    label: '流程运行',
    to: '/execution'
  },
  {
    label: '组织与权限',
    children: [
      {
        label: '用户管理',
        to: '/org/users'
      },
      {
        label: '角色管理',
        to: '/org/roles'
      }
    ]
  },
  {
    label: '系统管理',
    to: '/system'
  }
])

const items = computed(() => user.value ? consoleItems.value : publicItems.value)

const menus = computed(() => [
  [
    {
      label: user.value?.name,
      avatar: {
        loading: 'lazy',
        alt: user.value?.name,
        chip: {
          inset: true
        }
      },
      type: 'label'
    }
  ],
  [
    {
      label: '个人中心',
      icon: 'i-lucide-user',
      onSelect: () => navigateTo('/profile')
    },
    {
      label: '控制台首页',
      icon: 'i-lucide-layout-dashboard',
      onSelect: () => navigateTo('/dashboard')
    }
  ],
  [
    {
      label: '退出登录',
      icon: 'i-lucide-log-out',
      kbds: ['shift', 'meta', 'q'],
      onSelect: () => {
        logout()
        navigateTo('/')
      }
    }
  ]
])
</script>

<template>
  <UHeader>
    <template #left>
      <NuxtLink :to="user ? '/dashboard' : '/'">
        <AppLogo class="w-auto h-12 shrink-0" />
      </NuxtLink>
    </template>

    <UNavigationMenu
      :items="items"
      variant="link"
    />

    <template #right>
      <UColorModeButton />

      <UButton
        v-if="!user"
        icon="i-lucide-log-in"
        color="neutral"
        variant="ghost"
        to="/login"
        class="lg:hidden"
      />

      <UButton
        v-if="!user"
        label="登录"
        color="neutral"
        variant="outline"
        to="/login"
        class="hidden lg:inline-flex"
      />

      <UButton
        v-if="!user"
        label="注册"
        color="neutral"
        trailing-icon="i-lucide-arrow-right"
        class="hidden lg:inline-flex"
        to="/signup"
      />

      <UDropdownMenu
        v-if="user"
        :items="menus"
      >
        <UButton
          icon="i-lucide-user"
          class="rounded-full"
        />
      </UDropdownMenu>
    </template>

    <template #body>
      <UNavigationMenu
        :items="items"
        orientation="vertical"
        class="-mx-2.5"
      />

      <USeparator
        v-if="!user"
        class="my-6"
      />

      <UButton
        v-if="!user"
        label="Sign in"
        color="neutral"
        variant="subtle"
        to="/login"
        block
        class="mb-3"
      />
      <UButton
        v-if="!user"
        label="Sign up"
        color="neutral"
        to="/signup"
        block
      />
    </template>
  </UHeader>
</template>
