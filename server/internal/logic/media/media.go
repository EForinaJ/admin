package media

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	dao_media "server/internal/type/media/dao"
	dto_media "server/internal/type/media/dto"
	utils_error "server/internal/utils/error"
	utils_file "server/internal/utils/file"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMedia struct{}

// Delete implements service.IMedia.
func (s *sMedia) Delete(ctx context.Context, ids []int64) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysMedia.Table()).
		WhereIn(dao.SysMedia.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	path, err := tx.Model(dao.SysMedia.Table()).
		WhereIn(dao.SysMedia.Columns().Id, ids).
		Fields(dao.SysMedia.Columns().Path).
		Array()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	// 删除本地文件
	for _, v := range path {
		err = gfile.Remove("." + v.String())
		if err != nil {
			return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
		}
	}
	return
}

// GetList implements service.IMedia.
func (s *sMedia) GetList(ctx context.Context, req *dto_media.Query) (total int, res []*dao_media.List, err error) {
	m := dao.SysMedia.Ctx(ctx).Page(req.Page, req.Limit).
		Order(dao.SysMedia.Columns().CreateTime + " desc")
	if req.MediaType != "" {
		types := utils_file.GetFileTypes(req.MediaType)
		m = m.WhereIn(dao.SysMedia.Columns().MediaType, types)
	}

	if req.Name != "" {
		m = m.WhereLike(dao.SysMedia.Columns().OrName, "%"+req.Name+"%")
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var medias []*entity.SysMedia
	err = m.Scan(&medias)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	res = make([]*dao_media.List, len(medias))
	for i, v := range medias {
		var media *dao_media.List
		err = gconv.Scan(v, &media)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		res[i] = media
	}
	return
}

func init() {
	service.RegisterMedia(&sMedia{})
}
