package site

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dao_site "server/internal/type/site/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// GetInfo implements service.ISite.
func (s *sSite) GetInfo(ctx context.Context) (res *dao_site.Detail, err error) {
	options, err := g.Redis().Get(ctx, "site")
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}
	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		return
	}
	var site dao_site.Detail
	baseSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.BaseSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	baseJson, err := gjson.DecodeToJson(baseSetting)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	site.Title = baseJson.Get("title").String()
	site.Domain = baseJson.Get("domain").String()
	site.Logo = baseJson.Get("logo").String()
	site.Icon = baseJson.Get("icon").String()
	site.Description = baseJson.Get("description").String()
	site.Address = baseJson.Get("address").String()
	site.ICP = baseJson.Get("icp").String()
	site.Copyright = baseJson.Get("copyright").String()

	fileSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	fileJson, err := gjson.DecodeToJson(fileSetting)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	site.FileSize = fileJson.Get("fileSize").Int()
	site.FileType = fileJson.Get("fileType").Strings()
	site.ImageSize = fileJson.Get("imageSize").Int()
	site.ImageType = fileJson.Get("imageType").Strings()

	withdrawSetting, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	withdrawJson, err := gjson.DecodeToJson(withdrawSetting)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	site.Symbol = withdrawJson.Get("symbol").String()

	res = &site
	err = g.Redis().SetEX(ctx, "site", site, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return
}
