import request from '@/utils/http'

export function fetchGetAftersalesList(params: Aftersales.Params.Query) {
  return request.get<Aftersales.Response.List>({
    url: '/aftersales/list',
    params
  })
}
export function fetchGetAftersalesDetail(params: {id:number}) {
  return request.get<Aftersales.Response.Detail>({
    url: '/aftersales/detail',
    params
  })
}