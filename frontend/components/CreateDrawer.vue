<script lang="ts" setup>
import type { FormInst } from 'naive-ui'

const active = useState('create-drawer', () => false)

const formRef = ref<FormInst | null>(null)
const formValue = ref(
  {
    name: '',
    os: ''
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

function handleValidateClick (e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate(() => {

  })
}
</script>
<template>
  <lazy-n-drawer v-model:show="active" :width="502" placement="right" class="bg-zinc-950">
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
          <n-button @click="handleValidateClick">
            Create
          </n-button>
        </n-form-item>
      </n-form>
    </n-drawer-content>
  </lazy-n-drawer>
</template>
