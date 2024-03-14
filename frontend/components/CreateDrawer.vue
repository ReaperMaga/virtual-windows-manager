<script lang="ts" setup>
import { useQueryClient } from '@tanstack/vue-query'
import type { FormInst } from 'naive-ui'
import { api } from '~/composables/useQuery'
import { createVirtualWindowsSchema, type CreateVirtualWindow, virtualWindowsSchema } from '~/schemas'

const active = useState('create-drawer', () => false)

const formRef = ref<FormInst | null>(null)
const createLoading = ref<boolean>(false)
const formValue = ref<CreateVirtualWindow>(
  {
    name: '',
    os: undefined
  }
)
const osOptions = ['win7', 'win8', 'win10', 'win11'].map(
  v => ({
    label: v,
    value: v
  })
)

const rules = {
  name: {
    required: true,
    message: 'Please input your name',
    trigger: 'blur'
  },
  os: {
    required: true,
    trigger: ['blur', 'change'],
    message: 'Please select an os'
  }
}

const query = useQueryClient()

const message = useMessage()

const isMobileScreen = useIsMobileScreen()

function handleValidateClick (e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate((errors) => {
    if (!errors) {
      createLoading.value = true
      api({
        path: '/vws',
        method: 'POST',
        body: formValue.value
      }, virtualWindowsSchema, createVirtualWindowsSchema).then(() => {
        query.invalidateQueries({ queryKey: ['vws', 'list'] })
        active.value = false
        createLoading.value = false
        message.success('You successfully created a new vw: ' + formValue.value.name)
        formValue.value = { name: '', os: undefined }
      })
    }
  })
}
</script>
<template>
  <lazy-n-drawer
    v-model:show="active"
    :default-width="isMobileScreen ? '85vw' : '30vw'"
    placement="right"
    class="bg-zinc-950"
  >
    <n-drawer-content title="Create new virtual windows">
      <n-form
        ref="formRef"
        :label-width="80"
        :model="formValue"
        :rules="rules"
        size="large"
      >
        <n-form-item label="Name" path="name">
          <n-input v-model:value="formValue.name" placeholder="..." />
        </n-form-item>
        <n-form-item label="Select" path="os">
          <n-select
            v-model:value="formValue.os"
            placeholder="..."
            :options="osOptions"
          />
        </n-form-item>
        <n-form-item>
          <n-button :loading="createLoading" @click="handleValidateClick">
            Create
          </n-button>
        </n-form-item>
      </n-form>
    </n-drawer-content>
  </lazy-n-drawer>
</template>
