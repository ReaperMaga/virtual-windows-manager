<script lang="ts" setup>
import { useQueryClient } from '@tanstack/vue-query'
import { virtualWindowsSchema } from '~/schemas'

definePageMeta({
  layout: 'panel'
})

const { vwsQuery } = useQuery()

const { data, isLoading } = vwsQuery.list()

const createDrawerActive = useState('create-drawer', () => false)

const dialog = useDialog()

const message = useMessage()
const query = useQueryClient()

function handleDelete (id: string, name: string) {
  dialog.warning({
    title: 'Confirm',
    content: 'Are you sure?',
    positiveText: 'Sure',
    negativeText: 'Not Sure',
    onPositiveClick: () => {
      api({
        path: 'http://localhost:8080/vws/' + id,
        method: 'DELETE'
      }, virtualWindowsSchema).then(() => {
        query.invalidateQueries({ queryKey: ['vws', 'list'] })
        message.warning('You deleted ' + name)
      })
    }
  })
}

</script>

<template>
  <div v-if="!isLoading">
    <div v-if="data && data.length > 0" class="px-16 py-16 w-full grid grid-cols-3 gap-3">
      <n-card v-for="value in data" :key="value.id" title="Virtual Windows Manager" size="medium" class="transition hover:shadow-xl">
        <template #header>
          <div class="flex items-center gap-2">
            <span><Icon name="material-symbols:settings" size="25px" color="red" /></span>
            <h1 class="text-gray-300">
              {{ value.name }}
            </h1>
          </div>
        </template>
        <p class="text-gray-400 text-sm text-left">
          ID: {{ value.id }} <br>
          OS: {{ value.os }}
        </p>
        <div class="flex w-full">
          <div class="flex justify-end gap-3 w-full">
            <n-button type="error" ghost @click="handleDelete(value.id, value.name)">
              Delete
            </n-button>
            <n-button type="success" ghost>
              Start
            </n-button>
          </div>
        </div>
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
