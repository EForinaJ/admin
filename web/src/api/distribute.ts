import request from '@/utils/http'

export function fetchGetDistributeList(params: Distribute.Params.Query) {
  return request.get<Distribute.Response.List>({
    url: '/distribute/list',
    params
  })
}
export function fetchPostDistributeCreate(data: Distribute.Params.Model) {
  return request.post({
    url: '/distribute/create',
    data
  })
}

export function fetchPostDistributeCancel(data: Distribute.Params.Cancel) {
  return request.post({
    url: '/distribute/cancel',
    data
  })
}