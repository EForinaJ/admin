import request from '@/utils/http'

export function fetchGetCommentList(params: Comment.Params.Query) {
  return request.get<Comment.Response.List>({
    url: '/comment/list',
    params
  })
}

export function fetchPostCommentApply(data: Comment.Params.Apply) {
  return request.post({
    url: '/comment/apply',
    data
  })
}

export function fetchPostCommentDelete(data: {ids:number[]}) {
  return request.post({
    url: '/comment/delete',
    data
  })
}
