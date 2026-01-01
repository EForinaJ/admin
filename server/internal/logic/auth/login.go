package auth

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/lib/jwt"
	"server/internal/model/entity"
	dto_auth "server/internal/type/auth/dto"
	utils_error "server/internal/utils/error"
	utlis_lock "server/internal/utils/lock"
	"server/internal/utils/response"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Login implements admin_service.IAuth.
func (s *sAtuh) Login(ctx context.Context, req *dto_auth.Login) (res interface{}, err error) {
	obj := dao.SysManage.Ctx(ctx).Fields(
		dao.SysManage.Columns().Salt,
		dao.SysManage.Columns().Password,
		dao.SysManage.Columns().Id,
	)
	obj = obj.Where(dao.SysManage.Columns().Phone, req.Phone)

	userObj, err := obj.One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if userObj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	var user entity.SysManage
	userObj.Struct(&user)
	randPwd := consts.SYSTEMNAME + req.PassWord + user.Salt
	randPwd = gmd5.MustEncryptString(randPwd)
	token, err := jwt.GenToken(user.Id, user.Name)
	if err != nil {
		return "", utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
	}
	res = token
	if !strings.EqualFold(user.Password, randPwd) {
		// 设置密码错误次数
		errTimes, _ := utlis_lock.SetCount(ctx, consts.LoginCount+req.Phone,
			consts.LoginLock+req.Phone, 1800, 5)
		having := 5 - errTimes
		if having == 0 {
			_, err = g.Redis().Del(ctx, consts.LoginCount+req.Phone)
			if err != nil {
				return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
			}
			return nil, utils_error.Err(response.LOGIN_ERROR, "账号已锁定，请30分钟后再试")
		} else {
			return nil, utils_error.Err(response.LOGIN_ERROR, "密码不正确"+",还有"+gconv.String(having)+"次之后账号将锁定")
		}
	} else {

		// 判断是否是超管
		roleIds, err := dao.SysManageRole.Ctx(ctx).
			Where(dao.SysManageRole.Columns().ManageId, user.Id).
			Fields(dao.SysManageRole.Columns().RoleId).
			Array()
		if err != nil {
			return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		roleTypes, err := dao.SysRole.Ctx(ctx).
			WhereIn(dao.SysRole.Columns().Id, roleIds).
			Fields(dao.SysRole.Columns().Type).
			Array()
		if err != nil {
			return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		if !garray.NewIntArrayFrom(roleTypes.Ints()).Contains(consts.RoleTypeSuper) {
			return "", utils_error.Err(response.LOGIN_ERROR, "该职员不是超管，无法登录")
		}

		_, err = dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id, user.Id).
			Data(g.Map{
				dao.SysManage.Columns().LoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
				dao.SysManage.Columns().LoginTime: gtime.Now(),
			}).Update()
		if err != nil {
			return "", utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

		token, err := jwt.GenToken(user.Id, user.Name)
		if err != nil {
			return "", utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
		}
		res = token
	}
	return
}
