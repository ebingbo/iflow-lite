<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { ProcessAddBody } from '~/types/process'

const schema = z.object({
  code: z.string().min(1, '编码不能为空'),
  name: z.string().min(1, '名称不能为空'),
  description: z.string().optional()
})
const open = ref(false)
const isSubmitting = ref(false)

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  code: '',
  name: '',
  description: ''
})

const { addProcessWithFeedback } = useProcessActions()

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    isSubmitting.value = true
    await addProcessWithFeedback(event.data as ProcessAddBody, {
      successTitle: '成功',
      successDescription: `新的流程 ${event.data.name} 已添加`,
      errorTitle: '创建失败'
    })
    open.value = false
    state.code = ''
    state.name = ''
    state.description = ''
  } catch {
    return
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <UModal
    v-model:open="open"
    title="新建流程"
    description="添加新的流程定义"
  >
    <UButton
      label="新建流程"
      icon="i-lucide-plus"
    />

    <template #body>
      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormField
          label="名称"
          placeholder="请输入流程名称"
          name="name"
        >
          <UInput
            v-model="state.name"
            class="w-full"
          />
        </UFormField>
        <UFormField
          label="编码"
          placeholder="iss-xxx"
          name="code"
        >
          <UInput
            v-model="state.code"
            class="w-full"
          />
        </UFormField>

        <UFormField
          label="描述"
          placeholder="描述是什么样的流程"
          name="description"
        >
          <UTextarea
            v-model="state.description"
            class="w-full"
          />
        </UFormField>
        <div class="flex justify-end gap-2">
          <UButton
            label="取消"
            color="neutral"
            variant="subtle"
            @click="open = false"
          />
          <UButton
            label="创建"
            color="primary"
            variant="solid"
            type="submit"
            :loading="isSubmitting"
            :disabled="isSubmitting"
          />
        </div>
      </UForm>
    </template>
  </UModal>
</template>
