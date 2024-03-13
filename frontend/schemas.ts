import { z } from 'zod'
export const virtualWindowsSchema = z.object({
  id: z.string().min(1).max(64),
  name: z.string().max(32),
  os: z.enum(['win7', 'win8', 'win10', 'win11'])
})
