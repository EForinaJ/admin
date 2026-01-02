package service

import (
	"context"
	dao_prestore "server/internal/type/prestore/dao"
	dto_prestore "server/internal/type/prestore/dto"
)

// 定义显示接口
type IPrestore interface {
	// 检查余额是否足够
	Create(ctx context.Context, req *dto_prestore.Create) (err error)
	CheckApply(ctx context.Context, req *dto_prestore.Apply) (err error)
	GetDetail(ctx context.Context, id int64) (res *dao_prestore.Detail, err error)
	GetList(ctx context.Context, req *dto_prestore.Query) (total int, res []*dao_prestore.List, err error)
	Apply(ctx context.Context, req *dto_prestore.Apply) (err error)
}

// 定义接口变量
var localPrestore IPrestore

// 定义一个获取接口的方法
func Prestore() IPrestore {
	if localPrestore == nil {
		panic("implement not found for interface IPrestore, forgot register?")
	}
	return localPrestore
}

// 定义一个接口实现的注册方法
func RegisterPrestore(i IPrestore) {
	localPrestore = i
}
