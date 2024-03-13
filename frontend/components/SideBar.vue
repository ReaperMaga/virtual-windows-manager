<script setup lang="ts">

const props = defineProps({
  isNavbarCollapsed: Boolean
})

const { signOut } = useAuth()

const message = useMessage()

function logOut () {
  signOut({ callbackUrl: '/login' }).then(() => {
    message.warning('You have been signed out')
  })
}

</script>

<template>
  <div class="min-h-screen w-full shadow-lg glassy py-10 flex flex-col justify-between">
    <div class="flex flex-col">
      <h1 class="font-bold text-2xl text-gray-300 flex items-center justify-center gap-2">
        <span class="text-orange-500"><Icon name="icon-park-outline:computer" size="35px" /></span>
        <span v-if="!props.isNavbarCollapsed">Virtual VW</span>
      </h1>
    </div>
    <div class="flex items-center justify-center">
      <n-button v-if="!props.isNavbarCollapsed" type="error" ghost @click="logOut">
        Logout
      </n-button>
      <n-button
        v-if="props.isNavbarCollapsed"
        strong
        secondary
        circle
        type="error"
        @click="logOut"
      >
        <template #icon>
          <Icon name="material-symbols:logout-rounded" size="35px" />
        </template>
      </n-button>
    </div>
  </div>
</template>

<style>
.glassy {
  background:  rgba(0, 0, 0, 0.3);
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(5px);
  -webkit-backdrop-filter: blur(5px);
}
</style>
