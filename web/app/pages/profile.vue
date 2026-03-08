<script setup lang="ts">
definePageMeta({
  layout: 'default'
})

useSeoMeta({
  title: '个人中心',
  description: '查看和更新个人资料'
})

const toast = useToast()
const { user, fetchProfile, updateProfile, updatePassword, logout } = useAuth()

const profileForm = reactive({
  name: ''
})

const passwordForm = reactive({
  old_password: '',
  new_password: '',
  confirm_password: ''
})

const savingProfile = ref(false)
const savingPassword = ref(false)

if (!user.value) {
  await fetchProfile()
}

watchEffect(() => {
  profileForm.name = user.value?.name || ''
})

const saveProfile = async () => {
  const name = profileForm.name.trim()
  if (!name) {
    toast.add({
      title: '保存失败',
      description: '姓名不能为空',
      color: 'error'
    })
    return
  }
  try {
    savingProfile.value = true
    await updateProfile({ name })
    toast.add({
      title: '保存成功',
      description: '个人资料已更新',
      color: 'success'
    })
  } catch (error) {
    toast.add({
      title: '保存失败',
      description: error instanceof Error ? error.message : '更新资料失败',
      color: 'error'
    })
  } finally {
    savingProfile.value = false
  }
}

const savePassword = async () => {
  if (passwordForm.new_password.length < 8) {
    toast.add({
      title: '修改失败',
      description: '新密码至少 8 位',
      color: 'error'
    })
    return
  }
  if (passwordForm.new_password !== passwordForm.confirm_password) {
    toast.add({
      title: '修改失败',
      description: '两次输入的新密码不一致',
      color: 'error'
    })
    return
  }

  try {
    savingPassword.value = true
    await updatePassword({
      old_password: passwordForm.old_password,
      new_password: passwordForm.new_password
    })
    passwordForm.old_password = ''
    passwordForm.new_password = ''
    passwordForm.confirm_password = ''
    toast.add({
      title: '修改成功',
      description: '密码已更新，请重新登录',
      color: 'success'
    })
    logout()
    await navigateTo('/login')
  } catch (error) {
    toast.add({
      title: '修改失败',
      description: error instanceof Error ? error.message : '更新密码失败',
      color: 'error'
    })
  } finally {
    savingPassword.value = false
  }
}
</script>

<template>
  <UContainer class="space-y-4 pb-8">
    <UPageHeader
      title="个人中心"
      description="管理个人资料和账户安全"
    />

    <div class="grid gap-4 lg:grid-cols-2">
      <UCard>
        <template #header>
          <h3 class="font-semibold">
            基本信息
          </h3>
        </template>

        <div class="space-y-4">
          <UFormField label="邮箱">
            <UInput
              :model-value="user?.email || ''"
              disabled
              class="w-full"
            />
          </UFormField>

          <UFormField label="姓名">
            <UInput
              v-model="profileForm.name"
              class="w-full"
            />
          </UFormField>

          <UButton
            icon="i-lucide-save"
            color="primary"
            :loading="savingProfile"
            @click="saveProfile"
          >
            保存资料
          </UButton>
        </div>
      </UCard>

      <UCard>
        <template #header>
          <h3 class="font-semibold">
            修改密码
          </h3>
        </template>

        <div class="space-y-4">
          <UFormField label="旧密码">
            <UInput
              v-model="passwordForm.old_password"
              type="password"
              class="w-full"
            />
          </UFormField>

          <UFormField label="新密码">
            <UInput
              v-model="passwordForm.new_password"
              type="password"
              class="w-full"
            />
          </UFormField>

          <UFormField label="确认新密码">
            <UInput
              v-model="passwordForm.confirm_password"
              type="password"
              class="w-full"
            />
          </UFormField>

          <UButton
            icon="i-lucide-key-round"
            color="warning"
            :loading="savingPassword"
            @click="savePassword"
          >
            更新密码
          </UButton>
        </div>
      </UCard>
    </div>
  </UContainer>
</template>
