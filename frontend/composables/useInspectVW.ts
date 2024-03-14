import type { VirtualWindow } from '~/schemas'

export default () => {
  const active = useState('inspect-drawer', () => false)
  const vw = useState<VirtualWindow | undefined>('inspect-drawer-id', () => undefined)
  function open (vwObject: VirtualWindow) {
    vw.value = vwObject
    active.value = true
  }
  return { active, vw, open }
}
