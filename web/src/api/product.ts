import request from '@/utils/http'

export function fetchGetProductList(params: Product.Params.Query) {
    return request.get<Product.Response.List>({
      url: '/product/list',
      params
    })
}

export function fetchPostProductCreate(data: Product.Params.Model) {
  return request.post({
    url: '/product/create',
    data
  })
}

export function fetchGetProductEdit(params: {id:number}) {
  return request.get<Product.Response.Edit>({
    url: '/product/edit',
    params
  })
}

export function fetchPostProductEdit(data: Product.Params.Model) {
  return request.post({
    url: '/product/edit',
    data
  })
}

export function fetchPostProductDelete(data: {ids:number[]}) {
  return request.post({
    url: '/product/delete',
    data
  })
}