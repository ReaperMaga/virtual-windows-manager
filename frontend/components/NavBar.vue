<script lang="ts" setup>

const createDrawerActive = useState('create-drawer', () => false)

</script>
<template>
  <div class="h-16 glassy w-full flex justify-end items-center px-10">
    <n-button strong secondary circle type="primary" @click="createDrawerActive = true">
      <template #icon>
        <Icon name="material-symbols:add-rounded" size="35px" />
      </template>
    </n-button>
  </div>
</template>
