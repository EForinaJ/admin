import request from '@/utils/http'
// 获取用户列表
export function fetchGetUserList(params: User.Params.Query) {
    return request.get<User.Response.List>({
      url: '/user/list',
      params
    })
}

export function fetchPostUserCreate(data: User.Params.Model) {
  return request.post({
    url: '/user/create',
    data
  })
}

export function fetchGetUserEdit(params: {id:number}) {
  return request.get<User.Response.Edit>({
    url: '/user/edit',
    params
  })
}

export function fetchPostUserEdit(data: User.Params.Model) {
  return request.post({
    url: '/user/edit',
    data
  })
}

export function fetchPostUserDelete(data: {ids:number[]}) {
  return request.post({
    url: '/user/delete',
    data
  })
}


export function fetchGetUserDetail(params: {id:number}) {
  return request.get<User.Response.Detail>({
    url: '/user/detail',
    params
  })
}



export function fetchPostUserChangeBalance(data: User.Params.ChangeBalance) {
  return request.post({
    url: '/user/change/balance',
    data
  })
}



export function fetchPostUserRecharge(data: User.Params.Recharge) {
  return request.post({
    url: '/user/recharge',
    method: 'post',
    data:data
  });
}


export function fetchGetUserBalanceList(params: {
  id: number,
  page: number,
  limit: number
}) {
  return request.get<User.Response.BalanceList>({
    url: '/user/balance/list',
    params
  })
}

