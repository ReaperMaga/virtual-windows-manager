import { z } from 'zod'
import { useQuery } from '@tanstack/vue-query'
import { virtualWindowsSchema } from '~/schemas'

export default () => {
  return {
    vwsQuery: {
      list: () => useQuery({
        queryFn: () => api('http://localhost:8080/vws', z.array(virtualWindowsSchema)),
        queryKey: ['vws', 'list']
      })
    }
  }
}

async function api<T extends z.ZodTypeAny> (path: string, schema: T): Promise<z.infer<T>> {
  const { token } = useAuth()
  if (!token.value) {
    throw new Error('No token provided')
  }
  const data = await $fetch(path, {
    headers: {
      Authorization: token.value
    }
  })
  return schema.parse(data)
}
