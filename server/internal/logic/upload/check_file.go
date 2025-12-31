package upload

import (
	"context"
	"net/http"
	"server/internal/consts"
	"server/internal/dao"
	dao_system "server/internal/type/system/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckFile implements service.IUpload.
func (s *sUpload) CheckFile(ctx context.Context, file *ghttp.UploadFile) (err error) {
	fileSetting, err := dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().Key, consts.FileSetting).Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	fileJson, err := gjson.DecodeToJson(fileSetting)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var fileEntity *dao_system.File
	err = fileJson.Scan(&fileEntity)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	f, _ := file.Open()
	defer f.Close()

	buffer := make([]byte, 512)
	n, _ := f.Read(buffer)

	if isImgByContent(buffer[:n]) && isImgByExtension(file.Filename) {
		if file.Size > gconv.Int64(fileEntity.ImageSize)*1024*1024 {
			return utils_error.Err(response.UPDATE_FAILED, "图片已超出"+gconv.String(fileEntity.ImageSize)+"MB")
		}

		if !gstr.InArray(fileEntity.ImageType, gfile.Ext(file.Filename)) {
			return utils_error.Err(response.UPDATE_FAILED, "上传图片格式不正确")
		}
	} else {
		if file.Size > gconv.Int64(fileEntity.FileSize)*1024*1024 {
			return utils_error.Err(response.UPDATE_FAILED, "文件已超出"+gconv.String(fileEntity.ImageSize)+"MB")
		}

		if !gstr.InArray(fileEntity.FileType, gfile.Ext(file.Filename)) {
			return utils_error.Err(response.UPDATE_FAILED, "上传文件格式不正确")
		}
	}

	return
}

// IsImgByExtension 通过文件后缀名判断是否为图片
func isImgByExtension(filename string) bool {
	ext := gfile.Ext(filename) // 使用 GoFrame 的 gfile.Ext() 获取后缀
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".ico", ".webp":
		return true
	default:
		return false
	}
}

// isImgByContent 通过文件内容判断是否为图片
func isImgByContent(fileContent []byte) bool {
	// 只需读取前512字节即可判断类型[citation:6]
	contentType := http.DetectContentType(fileContent[:512])
	// 常见的图片MIME类型
	switch contentType {
	case "image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp":
		return true
	default:
		return false
	}
}
