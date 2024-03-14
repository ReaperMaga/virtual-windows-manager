import { z } from 'zod'
export const createVirtualWindowsSchema = z.object({
  name: z.string().max(32),
  os: z.enum(['win7', 'win8', 'win10', 'win11']).default('win8')

})

export const virtualWindowsSchema = createVirtualWindowsSchema.merge(z.object({
  id: z.string().min(1).max(64),
  running: z.boolean(),
  port: z.number()
}))

export type CreateVirtualWindow = z.input<typeof createVirtualWindowsSchema>
export type VirtualWindow = z.input<typeof virtualWindowsSchema>
