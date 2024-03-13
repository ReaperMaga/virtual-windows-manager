<script lang="ts" setup>
import {useAuth} from "#imports"
function loginToBackend() {
  const {signIn} = useAuth()
  console.log("Trying to login")
  signIn({username: "admin", password: "admin"}, {callbackUrl: "/"})
}
</script>

<template>
  <button @click="loginToBackend">Click me</button>
</template>
