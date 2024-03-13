// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@sidebase/nuxt-auth',
    '@nuxtjs/tailwindcss',
    '@bg-dev/nuxt-naiveui'
  ],
  typescript: {
    shim: false
  },
  auth: {
    baseURL: 'http://localhost:8080/auth',
    provider: {
      type: 'local',

      endpoints: {
        getSession: {path: '/session'},
        signOut: {
          path: '/logout', method: 'get'
        }
      },
      pages: {
        login: '/'
      },
      token: {
        signInResponseTokenPointer: '/sessionToken',
        type: "",
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
