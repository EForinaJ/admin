
import request from '@/utils/http'
/**
 * 获取用户信息
 * @returns 用户信息
 */
export function fetchGetAccountInfo() {
    return request.get<Account.Response.Info>({
      url: '/account'
      // 自定义请求头
      // headers: {
      //   'X-Custom-Header': 'your-custom-value'
      // }
    })
}


// 获取菜单列表
export function fetchGetAccountMenus() {
  return request.get<{
    list:Menu.Response.List[]
  }>({
    url: '/account/menus'
  })
}
