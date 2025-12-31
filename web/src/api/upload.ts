import request from '@/utils/http'

export function fetchUpload(data: FormData) {
    return request.post<{links:string[]}>({
      url: '/upload/file',
      data
    })
  }
  