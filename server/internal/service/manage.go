package service

import (
	"context"
	dao_manage "server/internal/type/manage/dao"
	dto_manage "server/internal/type/manage/dto"
)

// 定义显示接口
type IManage interface {

	// 获取列表
	GetList(ctx context.Context, req *dto_manage.Query) (total int, res []*dao_manage.List, err error)
	// 创建管理员
	Create(ctx context.Context, req *dto_manage.Create) (err error)
	// 获取修改信息
	GetEdit(ctx context.Context, id int64) (res *dao_manage.Edit, err error)
	// 修改内容
	Edit(ctx context.Context, req *dto_manage.Edit) (err error)
	// 删除
	Delete(ctx context.Context, ids []int64) (err error)

	CheckCreate(ctx context.Context, req *dto_manage.Create) (err error)
	CheckEdit(ctx context.Context, req *dto_manage.Edit) (err error)
}

// 定义接口变量
var localManage IManage

// 定义一个获取接口的方法
func Manage() IManage {
	if localManage == nil {
		panic("implement not found for interface IManage, forgot register?")
	}
	return localManage
}

// 定义一个接口实现的注册方法
func RegisterManage(i IManage) {
	localManage = i
}
