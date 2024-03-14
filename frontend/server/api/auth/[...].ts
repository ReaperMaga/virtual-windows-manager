import CredentialsProvider from 'next-auth/providers/credentials'
import { NuxtAuthHandler } from '#auth'

export default NuxtAuthHandler({
  // TODO: Replace this with a required secret for production
  secret: process.env.AUTH_SECRET || 'test-12345',
  jwt: {
    maxAge: 24 * 30 * 60 // 1 day
  },
  pages: {
    signIn: '/login'
  },
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    CredentialsProvider.default({
      name: 'Credentials',
      credentials: {
        username: { label: 'Username', type: 'text', placeholder: 'Username' },
        password: { label: 'Passwort', type: 'password', placeholder: 'Passwort' }
      },
      authorize: async (credentials: { username: string, password: string }) => {
        const response = await $fetch(process.env.NUXT_PUBLIC_API_BASE_URL + '/auth/login', {
          method: 'POST',
          body: JSON.stringify(credentials)
        })
        return response
      }
    })
  ],
  callbacks: {
    jwt: ({ token, user }) => {
      // user contains what you returned from the authorize function (id and name)
      if (user) {
        // Add informating into your JWT token (try to add as little as possible
        token.sessionToken = (user as any).sessionToken || ''
      }
      return token
    },
    session: async ({ session, token }) => {
      const response = await $fetch(process.env.NUXT_PUBLIC_API_BASE_URL + '/auth/session', {
        method: 'GET',
        headers: {
          Authorization: token.sessionToken as string
        }
      })
      if (!response) {
        throw new Error('Session expired')
      }
      return {
        ...session,
        sessionToken: token.sessionToken,
        user: response
      }
    }
  }
})
