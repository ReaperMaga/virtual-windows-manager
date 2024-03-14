export type DefaultSession = {
  user: { id: string, name: string, createdAt: string  }
  sessionToken: string
}

// Extend the NuxtAuth Session type with more information we pass in /server/api/auth/[...].ts
declare module 'next-auth' {
  /**
   * Returned by `useSession`, `getSession` and received as a prop on the `SessionProvider` React Context
   */
  type Session =  DefaultSession
}
