<script setup lang="ts">
import useIsSidebarCollapsed from '~/composeables/useIsSidebarCollapsed'

const route = useRoute()

const { status } = useAuth()
const isAuthenticated = computed(() => status.value === 'authenticated')

const isFooterExpanded = ref(false)
const isNavbarCollapsed = useIsSidebarCollapsed()
</script>

<template>
  <div class="flex">
    <n-layout position="absolute" has-sider>
      <n-layout-sider
        class="hidden md:block"
        :native-scrollbar="false"
        collapse-mode="width"
        :collapsed-width="75"
        :collapsed="isNavbarCollapsed"
        show-trigger="arrow-circle"
        bordered
        @collapse="isNavbarCollapsed = true"
        @expand="isNavbarCollapsed = false"
      >
        <SideBar :is-navbar-collapsed="isNavbarCollapsed" />
      </n-layout-sider>
      <n-layout-content>
        <div class="relative min-h-screen grow overflow-auto">
          <div class="flex flex-col w-full">
            <NavBar />
            <slot />
          </div>
        </div>
      </n-layout-content>
    </n-layout>
  </div>
</template>
