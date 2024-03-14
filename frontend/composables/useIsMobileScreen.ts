import { breakpointsTailwind, useBreakpoints } from '@vueuse/core'

export default () => useBreakpoints(breakpointsTailwind).smaller('md')
