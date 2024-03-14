import { z } from 'zod'
import { useQuery } from '@tanstack/vue-query'
import { virtualWindowsSchema } from '~/schemas'

export default () => {
  return {
    vwsQuery: {
      list: () => useQuery({
        queryFn: () => api({
          path: '/vws',
          method: 'GET'
        }, z.array(virtualWindowsSchema).nullish()),
        queryKey: ['vws', 'list']
      })
    }
  }
}

interface FetchOptions<I extends z.ZodTypeAny> {
  path: string
  method?: 'GET' | 'POST' | 'DELETE' | 'PATCH'
  body?: z.input<I>
}

export async function api<I extends z.ZodTypeAny, O extends z.ZodTypeAny> (
  options: FetchOptions<I>,
  schemaOutput: O,
  schemaInput?: I
): Promise<z.infer<O>> {
  const runtimeConfig = useRuntimeConfig()
  const { token } = useAuth()
  if (!token.value) {
    throw new Error('No token provided')
  }
  const data = await $fetch(runtimeConfig.public.apiBaseUrl + options.path, {
    method: options.method,
    body: schemaInput ? schemaInput?.parse(options.body) : undefined,
    headers: {
      Authorization: token.value
    }
  })
  return schemaOutput.parse(data)
}
