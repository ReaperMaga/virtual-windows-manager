import { z } from 'zod'
import { useQuery } from '@tanstack/vue-query'
import { virtualWindowsSchema } from '~/schemas'

export default () => {
  const { data } = useAuth()
  return {
    vwsQuery: {
      list: () => useQuery({
        queryFn: () => api({
          path: '/vws',
          method: 'GET'
        }, z.array(virtualWindowsSchema).nullish()),
        queryKey: ['vws', 'list'],
        enabled: () => !!data.value?.sessionToken
      })
    }
  }
}

interface FetchOptions<I extends z.ZodTypeAny> {
  path: string
  method?: 'GET' | 'POST' | 'DELETE' | 'PATCH'
  body?: z.input<I>
}

/**
 * This function is used to make API requests to the backend. It uses the fetch function to make the request and the zod library to validate the input and output of the request.
 * @param options - The options for the request
 * @param schemaOutput - The schema for the output of the request
 * @param schemaInput - The schema for the input of the request
 * @returns The output of the request
 */
export async function api<I extends z.ZodTypeAny, O extends z.ZodTypeAny> (
  options: FetchOptions<I>,
  schemaOutput: O,
  schemaInput?: I
): Promise<z.infer<O>> {
  const runtimeConfig = useRuntimeConfig()
  const { data: session } = useAuth()
  if (!session.value?.sessionToken) {
    throw new Error('No token provided')
  }
  const data = await $fetch(runtimeConfig.public.apiBaseUrl + options.path, {
    method: options.method,
    body: schemaInput ? schemaInput?.parse(options.body) : undefined,
    headers: {
      Authorization: session.value?.sessionToken
    }
  })
  return schemaOutput.parse(data)
}
