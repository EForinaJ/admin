package upload

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// MiniFile implements service.IUpload.
func (s *sUpload) MiniFile(ctx context.Context, file *ghttp.UploadFile) (res []string, err error) {
	fileConfig, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	fjson := gjson.New(fileConfig)
	path := fjson.Get("path").String()

	filePath := ".././public/" + path + "/" + gtime.Date()
	orName := file.Filename
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	file.Filename = name + gfile.Ext(file.Filename)
	_, err = file.Save(filePath, false)
	if err != nil {
		return nil, utils_error.Err(response.FILE_SAVE_ERROR, response.CodeMsg(response.FILE_SAVE_ERROR))
	}
	entites := &do.SysMedia{
		Path:       filePath[1:] + "/" + file.Filename,
		Ext:        gfile.Ext(file.Filename),
		MediaType:  gfile.Ext(file.Filename)[1:],
		Name:       name + gfile.Ext(file.Filename),
		CreateTime: gtime.Now(),
		OrName:     orName,
		Size:       gconv.String(file.Size),
	}

	_, err = dao.SysMedia.Ctx(ctx).Data(&entites).Insert()
	if err != nil {
		return nil, utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	res = []string{
		gconv.String(entites.Path),
	}
	return
}
