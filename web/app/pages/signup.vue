<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

definePageMeta({
  layout: 'auth'
})

useSeoMeta({
  title: 'Sign up',
  description: 'Create an account to get started'
})

const toast = useToast()

const fields = [{
  name: 'name',
  type: 'text' as const,
  label: '姓名',
  placeholder: '输入姓名',
  required: true
}, {
  name: 'email',
  type: 'text' as const,
  label: '邮箱',
  placeholder: '输入邮箱',
  required: true
}, {
  name: 'password',
  label: '密码',
  type: 'password' as const,
  placeholder: '输入密码',
  required: true
}]

const providers = [{
  label: 'Google',
  icon: 'i-simple-icons-google',
  onClick: () => {
    toast.add({
      title: 'Google',
      description: 'Login with Google'
    })
  }
}, {
  label: 'GitHub',
  icon: 'i-simple-icons-github',
  onClick: () => {
    toast.add({
      title: 'GitHub',
      description: 'Login with GitHub'
    })
  }
}]

const schema = z.object({
  name: z.string().min(1, '名称必填'),
  email: z.email('非法邮箱'),
  password: z.string('密码必填').min(8, '至少8个字符')
})

type Schema = z.output<typeof schema>

const { signup } = useAuth()

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    await signup(event.data.name, event.data.email, event.data.password)
    toast.add({
      title: '注册成功',
      description: '开始登录',
      color: 'success'
    })
    await navigateTo('/login')
  } catch (error) {
    toast.add({
      title: '注册失败',
      description: error instanceof Error ? error.message : '用户名或密码错误',
      color: 'error'
    })
  } finally {
  }
}
</script>

<template>
  <UAuthForm
    :fields="fields"
    :schema="schema"
    separator="或"
    title="创建账号"
    icon="i-lucide-user"
    :submit="{ label: '注册' }"
    @submit="onSubmit"
  >
    <template #description>
      已有账号？
      <ULink
        to="/login"
        class="text-primary font-medium"
      >登录
      </ULink>
      .
    </template>

    <template #footer>
      注册即表示您同意我们的
      <ULink
        to="/"
        class="text-primary font-medium"
      >服务条款
      </ULink>
      .
    </template>
  </UAuthForm>
</template>
