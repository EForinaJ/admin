import request from '@/utils/http'

// 更新基本信息
export function fetchGetSystemBase() {
    return request.get<System.Response.Base>({
      url: '/system/base',
    })
}
export function fetchPostSystemBase(data: System.Params.Base) {
  return request.post({
    url: '/system/base',
    data
  })
}

// 更新用户信息
export function fetchGetSystemUser() {
  return request.get<System.Response.User>({
    url: '/system/user',
  })
}
export function fetchPostSystemUser(data: System.Params.User) {
  return request.post({
    url: '/system/user',
    data
  })
}

// 更新提现信息
export function fetchGetSystemWithdraw() {
  return request.get<System.Response.Withdraw>({
    url: '/system/withdraw',
  })
}
export function fetchPostSystemWithdraw(data: System.Params.Withdraw) {
  return request.post({
    url: '/system/withdraw',
    data
  })
}

// 更新上传文件信息
export function fetchGetSystemFile() {
  return request.get<System.Response.File>({
    url: '/system/file',
  })
}
export function fetchPostSystemFile(data: System.Params.File) {
  return request.post({
    url: '/system/file',
    data
  })
}

// 更新微信小程序配置
export function fetchGetSystemWechatMiniProgram() {
  return request.get<System.Response.WechatMiniProgram>({
    url: '/system/wechat/mini/program',
  })
}
export function fetchPostSystemWechatMiniProgram(data: System.Params.WechatMiniProgram) {
  return request.post({
    url: '/system/wechat/mini/program',
    data
  })
}

// 更新微信支付配置
export function fetchGetSystemWechatPay() {
  return request.get<System.Response.WechatPay>({
    url: '/system/wechat/pay',
  })
}
export function fetchPostSystemWechatPay(data: System.Params.WechatPay) {
  return request.post({
    url: '/system/wechat/pay',
    data
  })
}


// =================================  协议政策
export function fetchGetSystemUserAgreement() {
  return request.get<{
    content: string
  }>({
    url: '/system/user/agreement',
  })
}
export function fetchPostSystemUserAgreement(data: {content: string}) {
  return request.post({
    url: '/system/user/agreement',
    data
  })
}

export function fetchGetSystemPrivacyPolicy() {
  return request.get<{
    content: string
  }>({
    url: '/system/privacy/policy',
  })
}
export function fetchPostSystemPrivacyPolicy(data: {content: string}) {
  return request.post({
    url: '/system/privacy/policy',
    data
  })
}

export function fetchGetSystemAboutUs() {
  return request.get<{
    content: string
  }>({
    url: '/system/about/us',
  })
}
export function fetchPostSystemAboutUs(data: {content: string}) {
  return request.post({
    url: '/system/about/us',
    data
  })
}