package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveFile(ctx context.Context, req *v1.SaveFileReq) (res *v1.SaveFileRes, err error) {
	err = service.System().SaveConfig(ctx, consts.FileSetting, "文件设置", req.File)
	return
}
