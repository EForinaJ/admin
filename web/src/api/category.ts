import request from '@/utils/http'

export function fetchGetCategoryList(params: Category.Params.Query) {
    return request.get<Category.Response.List>({
      url: '/category/list',
      params
    })
}

export function fetchPostCategoryCreate(data: Category.Params.Model) {
  return request.post({
    url: '/category/create',
    data
  })
}

export function fetchGetCategoryEdit(params: {id:number}) {
  return request.get<Category.Response.Edit>({
    url: '/category/edit',
    params
  })
}

export function fetchPostCategoryEdit(data: Category.Params.Model) {
  return request.post({
    url: '/category/edit',
    data
  })
}

export function fetchPostCategoryDelete(data: {ids:number[]}) {
  return request.post({
    url: '/category/delete',
    data
  })
}
