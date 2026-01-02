package service

import (
	"context"
	dao_user "server/internal/type/user/dao"
	dto_user "server/internal/type/user/dto"
)

// 定义显示接口
type IUser interface {
	CheckCreate(ctx context.Context, req *dto_user.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_user.Edit) (err error)
	CheckBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_user.Edit, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_user.Detail, err error)
	GetList(ctx context.Context, req *dto_user.Query) (total int, res []*dao_user.List, err error)

	ChangeBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error)
	Edit(ctx context.Context, req *dto_user.Edit) (err error)
	Create(ctx context.Context, req *dto_user.Create) (err error)

	GetBalanceList(ctx context.Context, req *dto_user.BalanceQuery) (total int, res []*dao_user.BalanceList, err error)
	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localUser IUser

// 定义一个获取接口的方法
func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// 定义一个接口实现的注册方法
func RegisterUser(i IUser) {
	localUser = i
}
