package service

import (
	"context"
	dao_comment "server/internal/type/cooment/dao"
	dto_comment "server/internal/type/cooment/dto"
)

// 定义显示接口
type IComment interface {
	// 检查余额是否足够
	Delete(ctx context.Context, ids []int64) (err error)
	CheckApply(ctx context.Context, req *dto_comment.Apply) (err error)
	GetList(ctx context.Context, req *dto_comment.Query) (total int, res []*dao_comment.List, err error)
	Apply(ctx context.Context, req *dto_comment.Apply) (err error)
}

// 定义接口变量
var localComment IComment

// 定义一个获取接口的方法
func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

// 定义一个接口实现的注册方法
func RegisterComment(i IComment) {
	localComment = i
}
