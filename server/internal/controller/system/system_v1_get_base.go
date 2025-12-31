package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func (c *ControllerV1) GetBase(ctx context.Context, req *v1.GetBaseReq) (res *v1.GetBaseRes, err error) {

	options, err := service.System().GetOne(ctx, consts.BaseSetting)
	if err != nil {
		return nil, err
	}

	json, err := gjson.DecodeToJson(options)
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	err = json.Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	return
}
