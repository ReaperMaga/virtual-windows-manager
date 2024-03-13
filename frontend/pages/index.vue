<script lang="ts" setup>
import { useAuth } from '#imports'

definePageMeta({
  layout: 'panel'
})

const { token, getSession } = useAuth()

const session = await getSession()

console.log(token.value)

</script>

<template>
</template>
