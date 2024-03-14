<script lang="ts" setup>
import { z } from 'zod'

const { active, vw } = useInspectVW()

const logs = ref('')

onMounted(() => {
  setInterval(() => {
    if (active.value) {
      api({
        path: '/vws/' + vw.value?.id + '/logs',
        method: 'GET'
      }, z.string()).then((val) => {
        logs.value = val
      })
    }
  }, 3000)
})

</script>
<template>
  <lazy-n-drawer v-model:show="active" :width="800" :max-width="1000" placement="right" class="bg-zinc-950">
    <n-drawer-content v-if="vw" :title="`Inspecting ${vw.name}`">
      <pre class="max-w-full whitespace-pre-line w-full bg-zinc-900 text-[0.7rem] shadow rounded-md px-7 py-2">{{ logs }}</pre>
    </n-drawer-content>
  </lazy-n-drawer>
</template>
