<script setup lang="ts">
const route = useRoute()

const items = computed(() => [
  {
    label: '流程',
    to: '/process'
  },
  {
    label: 'Docs',
    to: '/docs',
    active: route.path.startsWith('/docs')
  }, {
    label: 'Pricing',
    to: '/pricing'
  }, {
    label: 'Blog',
    to: '/blog'
  }, {
    label: 'Changelog',
    to: '/changelog'
  }])
const { user, logout } = useAuth()

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
      label: 'Profile',
      icon: 'i-lucide-user'
    }
  ],
  [
    {
      label: 'Logout',
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
      <NuxtLink to="/">
        <AppLogo class="w-auto h-6 shrink-0" />
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
