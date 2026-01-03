import request from '@/utils/http'

export function fetchGetDistributeList(params: Distribute.Params.Query) {
  return request.get<Distribute.Response.List>({
    url: '/distribute/list',
    params
  })
}
export function fetchGetDistributeDetail(params: {id:number}) {
  return request.get<Distribute.Response.Detail>({
    url: '/distribute/detail',
    params
  })
}
export function fetchPostDistributeCancel(data:Distribute.Params.Cancel) {
  return request.post({
    url: '/distribute/cancel',
    data
  })
}