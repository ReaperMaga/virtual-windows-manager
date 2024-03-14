<script lang="ts" setup>

definePageMeta({
  layout: 'panel'
})

const { vwsQuery } = useQuery()

const { data, isLoading } = vwsQuery.list()

const createDrawerActive = useState('create-drawer', () => false)

</script>

<template>
  <div v-if="!isLoading">
    <div v-if="data && data.length > 0" class="px-16 py-16 w-full grid grid-cols-3 gap-3">
      <n-card v-for="value in data" :key="value.id" title="Virtual Windows Manager" size="medium">
        <h1>{{ value.name }}</h1>
      </n-card>
    </div>
    <n-result
      v-else
      class="mt-20"
      status="404"
      title="It's pretty empty here"
      description="Why don't you add your first virtual windows machine here?"
    >
      <template #footer>
        <n-button @click="createDrawerActive = true">
          Create a new machine
        </n-button>
      </template>
    </n-result>
  </div>
</template>
