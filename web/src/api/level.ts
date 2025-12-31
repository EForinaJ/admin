import request from '@/utils/http'

export function fetchGetLevelList(params: Level.Params.Query) {
    return request.get<Level.Response.List>({
      url: '/level/list',
      params
    })
}

export function fetchPostLevelCreate(data: Level.Params.Model) {
  return request.post({
    url: '/level/create',
    data
  })
}

export function fetchGetLevelEdit(params: {id:number}) {
  return request.get<Level.Response.Edit>({
    url: '/level/edit',
    params
  })
}

export function fetchPostLevelEdit(data: Level.Params.Model) {
  return request.post({
    url: '/level/edit',
    data
  })
}

export function fetchPostLevelDelete(data: {ids:number[]}) {
  return request.post({
    url: '/level/delete',
    data
  })
}
