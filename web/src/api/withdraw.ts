import request from '@/utils/http'

export function fetchGetWithdrawList(params: Withdraw.Params.Query) {
  return request.get<Withdraw.Response.List>({
    url: '/withdraw/list',
    params
  })
}

export function fetchGetWithdrawDetail(params: {id:number}) {
  return request.get<Withdraw.Response.Detail>({
    url: '/withdraw/detail',
    params
  })
}

export function fetchPostWithdrawApply(data: Withdraw.Params.Apply) {
  return request.post({
    url: '/withdraw/apply',
    data
  })
}