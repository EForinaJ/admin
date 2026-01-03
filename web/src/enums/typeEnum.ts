
export enum RoleType {
  Super = 1, // 超管
  Kefu = 2, // 客服
}

export enum SexType {
  Male = 1,      // 男
  Female = 2,    // 女
  Other = 3,     // 其他
}

export enum ChangeBalanceType {
  UserChangeBalanceTypeSystem = 1,      // 系统变更
  UserChangeBalanceTypeRecharge = 2,      // 用户充值
  UserChangeBalanceTypePaidOrder = 3,      // 订单消费
}

export enum ChangeCommissionType {
  WitkeyChangeBalanceTypeSystem = 1,      // 系统变更
  WitkeyChangeBalanceTypeSettlement = 2,      // 报单结算
  WitkeyChangeBalanceTypeWithdraw  = 3,      // 提现佣金
}

export enum OrderLogType {
  Create           = 1,  // 创建订单
	AddDiscount      = 2,  // 添加优惠金额
  Cancel           = 3,  // 关闭订单
	Complete         = 4,  // 完成订单
	Paid             = 5,  // 确认收款订单
	AfterSales       = 6,  // 添加售后工单
	Distribute       = 7,  // 派发威客
	DistributeCancel = 8,  // 取消派单服务
	Start            = 9, // 开始服务
	Refund           = 10, // 订单手动退款
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

export enum DistributeType {
  Team = 1,      // 自带队伍
  Self = 2,    // 个人接单
}


