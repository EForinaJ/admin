import request from '@/utils/http'


/**
 * 
 * @param data 
 * @returns 
 */
export function fetchMenuCreate(data: Menu.Params.Model) {
  return request.post({
    url: '/menu/create',
    data
  })
}
  
/**
 * 
 * @returns 返回菜单列表
 */
export function fetchMenuAll() {
  return request.get<{
    list:Menu.Response.List[]
  }>({
    url: '/menu/all'
  })
}

/**
 * 
 * @param params 获取编辑信息ID
 * @returns 编辑信息
 */
export function fetchMenuGetEdit(params: {id:number}) {
  return request.get<Menu.Response.Edit>({
    url: '/menu/edit',
    params
  })
}

/**
 * 
 * @param data 编辑菜单数据
 * @returns 
 */
export function fetchMenuEdit(data: Menu.Params.Model) {
  return request.post({
    url: '/menu/edit',
    data
  })
}

export function fetchMenuDelete(data: {ids:number[]}) {
  return request.post({
    url: '/menu/delete',
    data
  })
}
