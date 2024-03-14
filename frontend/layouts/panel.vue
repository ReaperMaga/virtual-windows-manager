<script setup lang="ts">

const isNavbarCollapsed = useIsSidebarCollapsed()
</script>

<template>
  <div class="flex">
    <TheDrawers />
    <n-layout position="absolute" has-sider>
      <n-layout-sider
        class="hidden md:block"
        :native-scrollbar="false"
        collapse-mode="width"
        :collapsed-width="75"
        :collapsed="isNavbarCollapsed"
        show-trigger="bar"
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
