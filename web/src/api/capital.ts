import request from '@/utils/http'

export function fetchGetCapitalList(params: Capital.Params.Query) {
  return request.get<Capital.Response.List>({
    url: '/capital/list',
    params
  })
}

export function fetchPostCapitalDelete(data: {ids:number[]}) {
  return request.post({
    url: '/capital/delete',
    data
  })
}
