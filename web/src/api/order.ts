import request from '@/utils/http'

export function fetchGetOrderList(params: Order.Params.Query) {
    return request.get<Order.Response.List>({
      url: '/order/list',
      params
    })
}

export function fetchPostOrderAddDiscount(data: Order.Params.AddDiscount) {
  return request.post({
    url: '/order/add/discount',
    data
  })
}

export function fetchPostOrderPaid(data:{id:number,payMode:number}) {
  return request.post({
    url: '/order/paid',
    data
  })
}
export function fetchPostOrderCancel(data:{id:number}) {
  return request.post({
    url: '/order/cancel',
    data
  })
}
export function fetchPostOrderStart(data:{id:number}) {
  return request.post({
    url: '/order/start',
    data
  })
}
export function fetchPostOrderComplete(data:{id:number}) {
  return request.post({
    url: '/order/complete',
    data
  })
}
export function fetchPostOrderRefund(data:Order.Params.Refund) {
  return request.post({
    url: '/order/refund',
    data
  })
}

export function fetchGetOrderDetail(params: {id:number}) {
  return request.get<Order.Response.Detail>({
    url: '/order/detail',
    params
  })
}

export function fetchPostOrderDelete(data: {ids:number[]}) {
  return request.post({
    url: '/order/delete',
    data
  })
}

export function fetchGetOrderWitkeyList(params: Witkey.Params.Query) {
  return request.get<Witkey.Response.List>({
    url: '/order/witkey/list',
    params
  })
}


export function fetchPostOrderDistribute(data:Order.Params.Distribute) {
  return request.post({
    url: '/order/distribute',
    data
  })
}
export function fetchGetOrderDistributeList(params: Order.Params.DistributeQuery) {
  return request.get<Order.Response.DistributeList>({
    url: '/order/distribute/list',
    params
  })
}
export function fetchPostOrderDistributeCancel(data:Order.Params.DistributeCancel) {
  return request.post({
    url: '/order/distribute/cancel',
    data
  })
}
export function fetchGetOrderLogList(params: Order.Params.LogQuery) {
  return request.get<Order.Response.LogList>({
    url: '/order/log/list',
    params
  })
}