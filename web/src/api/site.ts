import request from '@/utils/http'

export function fetchSite() {
    return request.get<Site.Response.Info>({
      url: '/site',
    })
}

// 获取所属用户选项列表
export function fetchGetUserOptionsList(params: {phone:string}) {
  return request.get<{
    list: {
      id:number,
      name:string
    }[]
  }>({
    url: '/site/user/options',
    params
  })
}


// 获取所属游戏选项列表
export function fetchGetGameOptionsList() {
  return request.get<{
    list: {
      id:number,
      name:string
    }[]
  }>({
    url: '/site/game/options',
  })
}


// 获取所属头衔选项列表
export function fetchGetCategoryOptionsList(params: {gameId:number}) {
  return request.get<{
    list: {
      id:number,
      name:string
    }[]
  }>({
    url: '/site/category/options',
    params
  })
}




// 获取所属头衔选项列表
export function fetchGetTitleOptionsList(params: {gameId:number}) {
  return request.get<{
    list: {
      id:number,
      name:string
    }[]
  }>({
    url: '/site/title/options',
    params
  })
}

