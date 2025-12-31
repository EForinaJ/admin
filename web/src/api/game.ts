import request from '@/utils/http'

export function fetchGetGameList(params: Game.Params.Query) {
    return request.get<Game.Response.List>({
      url: '/game/list',
      params
    })
}

export function fetchPostGameCreate(data: Game.Params.Model) {
  return request.post({
    url: '/game/create',
    data
  })
}

export function fetchGetGameEdit(params: {id:number}) {
  return request.get<Game.Response.Edit>({
    url: '/game/edit',
    params
  })
}

export function fetchPostGameEdit(data: Game.Params.Model) {
  return request.post({
    url: '/game/edit',
    data
  })
}

export function fetchPostGameDelete(data: {ids:number[]}) {
  return request.post({
    url: '/game/delete',
    data
  })
}
