// https://nuxt.com/docs/api/configuration/nuxt-config
// @ts-ignore
export default defineNuxtConfig({
  modules: [
    '@sidebase/nuxt-auth',
    '@nuxtjs/tailwindcss',
    '@bg-dev/nuxt-naiveui',
    'nuxt-icon',
    '@vueuse/nuxt'
  ],
  typescript: {
    shim: false
  },
  naiveui: {
    colorModePreference: 'dark'
  },
  runtimeConfig: {
    public: {
      apiBaseUrl: '',
      vncBaseUrl: ''
    }
  },
  auth: {
    provider: {
      type: 'authjs'
    },
    globalAppMiddleware: {
      isEnabled: true
    }
  },
  tailwindcss: { exposeConfig: true }

})
