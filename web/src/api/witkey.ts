import request from '@/utils/http'

export function fetchGetWitkeyList(params: Witkey.Params.Query) {
    return request.get<Witkey.Response.List>({
      url: '/witkey/list',
      params
    })
}

export function fetchPostWitkeyCreate(data: Witkey.Params.Model) {
  return request.post({
    url: '/witkey/create',
    data
  })
}


export function fetchGetWitkeyEdit(params: {id:number}) {
  return request.get<Witkey.Response.Edit>({
    url: '/witkey/edit',
    params
  })
}

export function fetchPostWitkeyEdit(data: Witkey.Params.Model) {
  return request.post({
    url: '/witkey/edit',
    data
  })
}

export function fetchPostWitkeyDelete(data: {ids:number[]}) {
  return request.post({
    url: '/witkey/delete',
    data
  })
}

export function fetchGetWitkeyDetail(params: {id:number}) {
  return request.get<Witkey.Response.Detail>({
    url: '/witkey/detail',
    params
  })
}

export function fetchPostWitkeyChangeCommission(data: Witkey.Params.ChangeCommission) {
  return request.post({
    url: '/witkey/change/commission',
    data
  })
}

export function fetchGetWitkeyCommissionList(params: {
  id: number,
  page: number,
  limit: number
}) {
  return request.get<Witkey.Response.CommissionList>({
    url: '/witkey/commission/list',
    params
  })
}