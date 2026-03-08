<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'

definePageMeta({
  layout: 'auth'
})

useSeoMeta({
  title: '忘记密码',
  description: '通过邮箱重置密码'
})

const toast = useToast()
const { forgotPassword } = useAuth()
const loading = ref(false)

const schema = z.object({
  email: z.email('请输入有效邮箱')
})

type Schema = z.output<typeof schema>

const fields = [
  {
    name: 'email',
    type: 'text' as const,
    label: '邮箱',
    placeholder: '请输入注册邮箱',
    required: true
  }
]

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    loading.value = true
    await forgotPassword({ email: event.data.email })
    toast.add({
      title: '请求已提交',
      description: '如邮箱有效，系统将发送重置指引',
      color: 'success'
    })
    await navigateTo('/login')
  } catch (error) {
    toast.add({
      title: '提交失败',
      description: error instanceof Error ? error.message : '请稍后重试',
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
    title="忘记密码"
    icon="i-lucide-mail"
    :loading="loading"
    :submit="{ label: '发送重置请求' }"
    @submit="onSubmit"
  >
    <template #description>
      想起密码了？
      <ULink
        to="/login"
        class="text-primary font-medium"
      >返回登录
      </ULink>
      .
    </template>
  </UAuthForm>
</template>
