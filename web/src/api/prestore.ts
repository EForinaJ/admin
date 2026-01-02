import request from '@/utils/http'

export function fetchGetPrestoreList(params: Prestore.Params.Query) {
  return request.get<Prestore.Response.List>({
    url: '/prestore/list',
    params
  })
}

export function fetchGetPrestoreDetail(params: {id:number}) {
  return request.get<Prestore.Response.Detail>({
    url: '/prestore/detail',
    params
  })
}

export function fetchPostPrestoreApply(data: Prestore.Params.Apply) {
  return request.post({
    url: '/prestore/apply',
    data
  })
}

export function fetchPostPrestoreCreate(data: Prestore.Params.Moadl) {
  return request.post({
    url: '/prestore/create',
    data
  })
}