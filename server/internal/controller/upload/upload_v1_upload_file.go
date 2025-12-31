package upload

import (
	"context"

	v1 "server/api/upload/v1"
	"server/internal/service"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	err = service.Upload().CheckFile(ctx, req.File)
	if err != nil {
		return nil, err
	}

	links, err := service.Upload().MiniFile(ctx, req.File)
	if err != nil {
		return nil, err
	}
	res = &v1.UploadFileRes{
		Links: links,
	}
	return
}
