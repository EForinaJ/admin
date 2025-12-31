export enum SexType {
  Male = 1,      // 男
  Female = 2,    // 女
  Other = 3,     // 其他
}

export enum ChangeBalanceType {
  UserChangeBalanceTypeSystem = 1,      // 系统变更
  UserChangeBalanceTypeRecharge = 2,      // 用户充值
}

export enum ChangeCommissionType {
  WitkeyChangeBalanceTypeSystem = 1,      // 系统变更
  WitkeyChangeBalanceTypeSettlement = 2,      // 报单结算
}


export enum CpitalType {
  PaymentOrder = 1,// 支付订单
  RefundOrder = 2,// 订单退款
  PaymentRecharge = 3,// 充值余额
  WithdrawCommission = 4,// 佣金提现
}

export enum AfterSalesType {
  ServiceNotCompleted = 1,      // 服务无法完成
  Cheater = 2,      // 威客作弊
  Other = 3,      // 其他
}

export enum WithdrawType {
  AlyPay = 1,      // 支付宝
  Wechat = 2,      // 微信
}


