export enum FundMode {
    Add = 1,      // 增加
    Sub = 2,    // 减少
}

export enum UserLogMode {
    DeleteWitkey = 8, // 删除威客
    EditWitkey = 7, // 修改威客
    AddWitkey = 6, // 添加威客
    Revoke = 5,// 撤回预存充值
    Balance = 4,// 变更余额
    Recharge = 3,// 预存充值
    Edit  = 2,// 修改内容
    Create = 1,// 创建用户
}


export enum PayMode {
    AlyPay = 1,// 支付宝
    Wechat = 2,// 微信
    Balance = 3,// 余额
    PersonalTransfer = 4,// 人工转账
}