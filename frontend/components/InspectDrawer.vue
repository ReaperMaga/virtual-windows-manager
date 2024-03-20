<script lang="ts" setup>
import { z } from 'zod'

const { active, vw } = useInspectVW()

const isMobileScreen = useIsMobileScreen()

const logs = ref('')

// This file is used to create a drawer that will be used to inspect a virtual windows machine
// It uses the setInterval function to get the logs of the virtual windows machine every 3 seconds

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
  <lazy-n-drawer
    v-model:show="active"
    :default-width="isMobileScreen ? '85vw' : '30vw'"
    placement="right"
    class="bg-zinc-950"
  >
    <n-drawer-content v-if="vw" :title="`Inspecting ${vw.name}`">
      <pre class="max-w-full whitespace-pre-line w-full bg-zinc-900 text-[0.7rem] shadow rounded-md px-7 py-2">{{ logs }}</pre>
    </n-drawer-content>
  </lazy-n-drawer>
</template>
