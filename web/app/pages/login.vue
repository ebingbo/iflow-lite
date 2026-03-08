<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

definePageMeta({
  layout: 'auth'
})

useSeoMeta({
  title: 'Login',
  description: 'Login to your account to continue'
})

const toast = useToast()

const fields = [{
  name: 'email',
  type: 'text' as const,
  label: '邮箱',
  placeholder: '输入邮箱',
  required: true
}, {
  name: 'password',
  label: '密码',
  type: 'password' as const,
  placeholder: '输入密码'
}, {
  name: 'remember',
  label: '记住我',
  type: 'checkbox' as const
}]

const schema = z.object({
  email: z.email('非法邮箱'),
  password: z.string('密码必填').min(8, '至少8个字符')
})

type Schema = z.output<typeof schema>
const loading = ref(false)
const { token, login } = useAuth()

if (token.value) {
  await navigateTo('/')
}

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    loading.value = true
    await login(event.data.email, event.data.password)
    toast.add({
      title: '登录成功',
      description: '欢迎回来',
      color: 'success'
    })
    await navigateTo('/dashboard')
  } catch (error) {
    toast.add({
      title: '登录失败',
      description: error instanceof Error ? error.message : '用户名或密码错误',
      color: 'error'
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <UAuthForm
    :fields="fields"
    :schema="schema"
    title="欢迎回来"
    icon="i-lucide-lock"
    :submit="{
      label: '登录'
    }"
    @submit="onSubmit"
  >
    <template #description>
      没有账号？
      <ULink
        to="/signup"
        class="text-primary font-medium"
      >注册
      </ULink>
      .
    </template>

    <template #password-hint>
      <ULink
        to="/forgot-password"
        class="text-primary font-medium"
        tabindex="-1"
      >忘记密码？
      </ULink>
    </template>

    <template #footer>
      登录即表示您同意我们的
      <ULink
        to="/"
        class="text-primary font-medium"
      >服务条款
      </ULink>
      .
    </template>
  </UAuthForm>
</template>
