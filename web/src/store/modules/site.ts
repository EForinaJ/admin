import { defineStore } from 'pinia'
// 表格
export const useSiteStore = defineStore(
  'siteStore',
  () => {
    
    // 站点信息
    const info = ref<Partial<Site.Response.Info>>({})

    // 计算属性：获取站点信息
    const getInfo = computed(() => info.value)
    /**
     * 设置站点信息
     * @param newInfo 新的站点信息
     */
    const setInfo = (newInfo: Site.Response.Info) => {
        info.value = newInfo
    }

    

    return {
        info,
        getInfo,
        setInfo
    }
  },
  {
    persist: {
      key: 'site',
      storage: localStorage
    }
  }
)
