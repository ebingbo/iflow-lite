<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import type { Process, ProcessUpdateBody } from '~/types/process'

const props = defineProps<{
  open: boolean
  process: Process | null
}>()

const emit = defineEmits<{
  (event: 'update:open', value: boolean): void
  (event: 'success'): void
}>()

const schema = z.object({
  name: z.string().min(1, '名称不能为空'),
  description: z.string().optional()
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  name: '',
  description: ''
})

const isSubmitting = ref(false)
const { updateProcessWithFeedback } = useProcessActions()

watch(
  () => [props.open, props.process] as const,
  ([open, process]) => {
    if (!open || !process) return
    state.name = process.name
    state.description = process.description || ''
  },
  { immediate: true }
)

function closeModal() {
  emit('update:open', false)
}

async function onSubmit(event: FormSubmitEvent<Schema>) {
  if (!props.process) return

  try {
    isSubmitting.value = true
    const payload: ProcessUpdateBody = {
      id: props.process.id,
      name: event.data.name,
      description: event.data.description || ''
    }
    await updateProcessWithFeedback(payload, {
      successTitle: '更新成功',
      successDescription: `流程 ${payload.name} 已更新`,
      errorTitle: '更新失败'
    })
    emit('success')
    closeModal()
  } catch {
    return
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <UModal
    :open="open"
    title="编辑流程"
    description="仅支持更新流程名称和描述"
    @update:open="emit('update:open', $event)"
  >
    <template #body>
      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormField
          label="名称"
          name="name"
        >
          <UInput
            v-model="state.name"
            class="w-full"
          />
        </UFormField>

        <UFormField
          label="描述"
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
            @click="closeModal"
          />
          <UButton
            label="保存"
            color="primary"
            type="submit"
            :loading="isSubmitting"
            :disabled="isSubmitting"
          />
        </div>
      </UForm>
    </template>
  </UModal>
</template>
