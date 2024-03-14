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
      apiBaseUrl: process.env.API_BASE_URL,
      vncBaseUrl: process.env.VNC_BASE_URL
    }
  },
  auth: {
    baseURL: process.env.API_BASE_URL + '/auth',
    provider: {
      type: 'local',
      endpoints: {
        getSession: { path: '/session' },
        signOut: {
          path: '/logout', method: 'get'
        }
      },
      pages: {
        login: '/login'
      },
      token: {
        signInResponseTokenPointer: '/sessionToken',
        type: '',
        maxAgeInSeconds: 360000
      },
      sessionDataType: {
        id: 'string',
        name: 'string'
      }
    },
    session: {
      enableRefreshOnWindowFocus: true,
      enableRefreshPeriodically: 5000
    },
    globalAppMiddleware: {
      isEnabled: true
    }
  }
})
