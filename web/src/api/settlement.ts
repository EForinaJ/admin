import request from '@/utils/http'

export function fetchGetSettlementList(params: Settlement.Params.Query) {
  return request.get<Settlement.Response.List>({
    url: '/settlement/list',
    params
  })
}

export function fetchGetSettlementDetail(params: {id:number}) {
  return request.get<Settlement.Response.Detail>({
    url: '/settlement/detail',
    params
  })
}

export function fetchPostSettlementApply(data: Settlement.Params.Apply) {
  return request.post({
    url: '/settlement/apply',
    data
  })
}