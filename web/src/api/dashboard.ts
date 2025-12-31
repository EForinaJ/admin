import request from '@/utils/http'

export function fetchGetDashboardDetail() {
  return request.get<Dashboard.Response.Detail>({
    url: '/dashboard/detail',
  })
}