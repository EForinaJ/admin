import request from '@/utils/http'

export function fetchGetRechargeList(params: Recharge.Params.Query) {
  return request.get<Recharge.Response.List>({
    url: '/recharge/list',
    params
  })
}

export function fetchPostRechargeDelete(data: {ids:number[]}) {
  return request.post({
    url: '/recharge/delete',
    data
  })
}
