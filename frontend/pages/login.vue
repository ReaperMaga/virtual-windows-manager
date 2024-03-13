<script lang="ts">
import { useMessage } from 'naive-ui'
import { useAuth } from '#imports'

export default defineComponent({
  setup () {
    const { signIn } = useAuth()
    const formRef = ref<null>(null)
    const message = useMessage()
    const formValue = ref(
      {
        user: {
          username: '',
          password: ''
        }
      }
    )
    return {
      formRef,
      size: ref<'small' | 'medium' | 'large'>('medium'),
      formValue,
      rules: {
        user: {
          name: {
            required: true,
            message: 'Please input your name',
            trigger: 'blur'
          },
          password: {
            required: true,
            message: 'Please input your password',
            trigger: ['input', 'blur']
          }
        }
      },
      handleValidateClick (e: MouseEvent) {
        e.preventDefault()
        formRef.value?.validate((errors) => {
          if (!errors) {
            signIn(formValue.value.user, { callbackUrl: '/' }).then(() => {
              message.success('Successfully logged in')
            })
          } else {
            message.error('Wrong credentials')
          }
        })
      }
    }
  }
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center w-1/4 mx-auto">
    <n-card title="Virtual Windows Manager" size="huge">
      <n-form
        ref="formRef"
        :label-width="80"
        :model="formValue"
        :rules="rules"
        :size="large"
      >
        <n-form-item label="Username" path="user.username">
          <n-input v-model:value="formValue.user.username" placeholder="..." />
        </n-form-item>
        <n-form-item label="Password" path="user.password">
          <n-input v-model:value="formValue.user.password" placeholder="..." type="password" />
        </n-form-item>
        <n-form-item>
          <n-button @click="handleValidateClick">
            Submit
          </n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </div>
</template>
