import request from '@/utils/http'


export function fetchRoleList(params: Role.Params.Query) {
    return request.get<Role.Response.List>({
      url: '/role/list',
      params
    })
}

export function fetchRoleCreate(data: Role.Params.Model) {
  return request.post({
    url: '/role/create',
    data
  })
}

export function fetchRoleGetEdit(params: {id:number}) {
  return request.get<Role.Response.Edit>({
    url: '/role/edit',
    params
  })
}

export function fetchRoleEdit(data: Role.Params.Model) {
  return request.post({
    url: '/role/edit',
    data
  })
}

export function fetchRoleDelete(data: {ids:number[]}) {
  return request.post({
    url: '/role/delete',
    data
  })
}



export function fetchRoleMenuAll() {
  return request.get<{
    list:Menu.Response.List[]
  }>({
    url: '/role/all/menu'
  })
}


export function fetchRoleMenus(params: {id:number}) {
  return request.get<{
    menus:number[]
  }>({
    url: '/role/menus',
    params
  })
}


export function fetchRoleEditMenus(data: {id:number,menus:number[]}) {
  return request.post({
    url: '/role/edit/menus',
    data
  })
}


export function fetchRolePermissionAll() {
  return request.get<{
    list:Permission.Response.Info[]
  }>({
    url: '/role/all/permission'
  })
}


export function fetchRolePermissions(params: {id:number}) {
  return request.get<{
    permissiones:number[]
  }>({
    url: '/role/permissions',
    params
  })
}


export function fetchRoleEditPermissions(data: {id:number,permissions:number[]}) {
  return request.post({
    url: '/role/edit/permissions',
    data
  })
}