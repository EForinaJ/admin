// hooks/useDebounce.ts
import { onUnmounted } from 'vue'

export function useDebounce() {
  let timer: NodeJS.Timeout | null = null

  const debounce = <T extends (...args: any[]) => any>(
    fn: T,
    delay: number = 500
  ): ((...args: Parameters<T>) => void) => {
    return (...args: Parameters<T>) => {
      if (timer) {
        clearTimeout(timer)
      }
      timer = setTimeout(() => {
        fn(...args)
      }, delay)
    }
  }

  // 取消防抖
  const cancelDebounce = () => {
    if (timer) {
      clearTimeout(timer)
      timer = null
    }
  }

  // 立即执行并取消防抖
  const flushDebounce = <T extends (...args: any[]) => any>(
    fn: T,
    ...args: Parameters<T>
  ) => {
    cancelDebounce()
    fn(...args)
  }

  // 清理
  onUnmounted(() => {
    cancelDebounce()
  })

  return {
    debounce,
    cancelDebounce,
    flushDebounce
  }
}