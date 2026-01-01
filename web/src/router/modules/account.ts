import { AppRouteRecord } from '@/types/router'

export const accountRoutes: AppRouteRecord = {
  path: '/account',
  name: 'Account',
  component: '/index/index',
  meta: {
    title: '账户管理',
    isHide: true,
    isHideTab: true,
    icon: 'solar:settings-minimalistic-bold',
  },
  children: [
    {
      path: 'center',
      name: 'AccountCenter',
      component: '/account/user-center',
      meta: {
        title: '账户中心',
        keepAlive: true,
        isHide: true,
        isHideTab: true,
        activePath: '/dashboard'
      }
    },
  ]
}
