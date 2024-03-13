<template>
  <div class="h-16 glassy w-full flex justify-end items-center px-10">
    <n-button strong secondary circle type="primary" @click="activate('right')">
      <template #icon>
        <Icon name="material-symbols:add-rounded" size="35px" />
      </template>
    </n-button>

    <n-drawer v-model:show="active" :width="502" :placement="placement" content-class="drawer-bg">
      <n-drawer-content title="Create new virtual windows" />
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
import type { DrawerPlacement } from 'naive-ui'

const active = ref(false)
const placement = ref<DrawerPlacement>('right')
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}
</script>
