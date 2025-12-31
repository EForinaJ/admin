import request from '@/utils/http'

export function fetchGetTitleList(params: Title.Params.Query) {
    return request.get<Title.Response.List>({
      url: '/title/list',
      params
    })
}

export function fetchPostTitleCreate(data: Title.Params.Model) {
  return request.post({
    url: '/title/create',
    data
  })
}

export function fetchGetTitleEdit(params: {id:number}) {
  return request.get<Title.Response.Edit>({
    url: '/title/edit',
    params
  })
}

export function fetchPostTitleEdit(data: Title.Params.Model) {
  return request.post({
    url: '/title/edit',
    data
  })
}

export function fetchPostTitleDelete(data: {ids:number[]}) {
  return request.post({
    url: '/title/delete',
    data
  })
}
