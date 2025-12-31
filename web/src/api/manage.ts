import request from '@/utils/http'


export function fetchManageList(params: Manage.Params.Query) {
    return request.get<Manage.Response.List>({
      url: '/manage/list',
      params
    })
}


export function fetchManageAllRole() {
  return request.get<{
    list:Role.Response.Info[]
  }>({
    url: '/manage/all/role',
  })
}


export function fetchManageCreate(data: Manage.Params.Model) {
  return request.post({
    url: '/manage/create',
    data
  })
}

export function fetchManageGetEdit(params: {id:number}) {
  return request.get<Manage.Response.Edit>({
    url: '/manage/edit',
    params
  })
}

export function fetchManageEdit(data: Manage.Params.Model) {
  return request.post({
    url: '/manage/edit',
    data
  })
}

export function fetchManageDelete(data: {ids:number[]}) {
  return request.post({
    url: '/manage/delete',
    data
  })
}