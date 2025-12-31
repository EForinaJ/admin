import request from '@/utils/http'

export function fetchPermissionCreate(data: Permission.Params.Model) {
    return request.post({
      url: '/permission/create',
      data
    })
}

/**
 * 
 * @returns 返回菜单列表
 */
export function fetchPermissionAll() {
  return request.get<{
    list:Permission.Response.Info[]
  }>({
    url: '/permission/all'
  })
}
  

export function fetchPermissionGetEdit(params: {id:number}) {
  return request.get<Permission.Response.Edit>({
    url: '/permission/edit',
    params
  })
}

/**
 * 
 * @param data 编辑菜单数据
 * @returns 
 */
export function fetchPermissionEdit(data: Permission.Params.Model) {
  return request.post({
    url: '/permission/edit',
    data
  })
}

export function fetchPermissionDelete(data: {ids:number[]}) {
  return request.post({
    url: '/permission/delete',
    data
  })
}
