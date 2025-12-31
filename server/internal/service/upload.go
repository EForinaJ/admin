package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 定义显示接口
type IUpload interface {
	// 小文件上传
	MiniFile(ctx context.Context, file *ghttp.UploadFile) (res []string, err error)
	CheckFile(ctx context.Context, file *ghttp.UploadFile) (err error)

	// // 获取切片标识
	// ChunkIdentifier(ctx context.Context, identifier string) (res []int, err error)
	// //  上传切片
	// UploadChunk(ctx context.Context, req *dto_upload.Chunk) (err error)
	// //  切片合并
	// MergeChunk(ctx context.Context, req *dto_upload.Merge) (res []string, err error)
}

// 定义接口变量
var localUpload IUpload

// 定义一个获取接口的方法
func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

// 定义一个接口实现的注册方法
func RegisterUpload(i IUpload) {
	localUpload = i
}
