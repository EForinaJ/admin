package service

import (
	"context"
	dao_role "server/internal/type/role/dao"
	dto_role "server/internal/type/role/dto"
)

// 定义显示接口
type IRole interface {
	// 获取全部角色
	GetAll(ctx context.Context) (res []*dao_role.List, err error)

	// 获取列表
	GetList(ctx context.Context, req *dto_role.Query) (total int, res []*dao_role.List, err error)
	// 创建
	Create(ctx context.Context, req *dto_role.Create) (err error)
	// 获取修改信息
	GetEdit(ctx context.Context, id int64) (res *dao_role.Edit, err error)
	// 修改内容
	Edit(ctx context.Context, req *dto_role.Edit) (err error)
	// 删除
	Delete(ctx context.Context, ids []int64) (err error)
	// 检查是否存在
	CheckCreate(ctx context.Context, req *dto_role.Create) (err error)
	//  排除ID当前是否存在
	CheckEdit(ctx context.Context, req *dto_role.Edit) (err error)
}

// 定义接口变量
var localRole IRole

// 定义一个获取接口的方法
func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

// 定义一个接口实现的注册方法
func RegisterRole(i IRole) {
	localRole = i
}
