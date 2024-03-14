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

const { open } = useInspectVW()

function handleDelete (id: string, name: string) {
  dialog.warning({
    title: 'Confirm',
    content: 'Are you sure?',
    positiveText: 'Sure',
    negativeText: 'Not Sure',
    onPositiveClick: () => {
      api({
        path: '/vws/' + id,
        method: 'DELETE'
      }, virtualWindowsSchema).then(() => {
        query.invalidateQueries({ queryKey: ['vws', 'list'] })
        message.warning('You deleted ' + name)
      })
    }
  })
}

function handleStart (id: string, name: string) {
  api({
    path: '/vws/' + id + '/start',
    method: 'POST'
  }, virtualWindowsSchema).then(() => {
    query.invalidateQueries({ queryKey: ['vws', 'list'] })
    message.warning('You started ' + name)
  })
}

function handleStop (id: string, name: string) {
  dialog.warning({
    title: 'Confirm',
    content: 'Are you sure?',
    positiveText: 'Sure',
    negativeText: 'Not Sure',
    onPositiveClick: () => {
      api({
        path: '/vws/' + id + '/stop',
        method: 'POST'
      }, virtualWindowsSchema).then(() => {
        query.invalidateQueries({ queryKey: ['vws', 'list'] })
        message.warning('You stopped ' + name)
      })
    }
  })
}

</script>

<template>
  <div v-if="!isLoading">
    <div v-if="data && data.length > 0" class="p-8 md:p-16 w-full grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3">
      <n-card v-for="value in data" :key="value.id" title="Virtual Windows Manager" size="medium" class="transition hover:shadow-xl">
        <template #header>
          <div class="flex items-center gap-2">
            <span v-if="!value.running"><Icon name="material-symbols:settings" size="25px" color="red" /></span>
            <span v-else><Icon name="material-symbols:settings" size="25px" color="green" class="animate-spin" /></span>
            <h1 class="text-gray-300">
              {{ value.name }}
            </h1>
          </div>
        </template>
        <p class="text-gray-400 text-sm text-left">
          ID: {{ value.id }} <br>
          OS: {{ value.os }}
        </p>
        <div class="flex w-full mt-6">
          <div v-if="!value.running" class="flex justify-end gap-3 w-full">
            <n-button type="error" ghost @click="handleDelete(value.id, value.name)">
              Delete
            </n-button>
            <n-button type="success" ghost @click="handleStart(value.id, value.name)">
              Start
            </n-button>
          </div>
          <div v-if="value.running" class="flex justify-end gap-3 w-full">
            <n-button type="info" ghost @click="open(value)">
              Inspect
            </n-button>
            <n-button type="warning" ghost @click="navigateTo('http://localhost:' + value.port, {open: {target: '_blank'}})">
              Open in VNC
            </n-button>
            <n-button type="error" ghost @click="handleStop(value.id, value.name)">
              Stop
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
