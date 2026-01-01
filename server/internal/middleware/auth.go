package middleware

import (
	"server/internal/lib/jwt"
	"server/internal/utils/response"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Auth(r *ghttp.Request) {
	// 获取token
	authHeader := r.GetHeader("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	mc, err := jwt.ParseToken(parts[1])
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	// 判断是否是超管
	// roleIds, err := dao.SysManageRole.Ctx(r.GetCtx()).
	// 	Where(dao.SysManageRole.Columns().ManageId, mc.Id).
	// 	Fields(dao.SysManageRole.Columns().RoleId).
	// 	Array()
	// if err != nil {
	// 	response.Error(r).SetCode(response.DB_READ_ERROR).
	// 		SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	// }

	// roleTypes, err := dao.SysRole.Ctx(r.GetCtx()).
	// 	WhereIn(dao.SysRole.Columns().Id, roleIds).
	// 	Fields(dao.SysRole.Columns().Type).
	// 	Array()
	// if err != nil {
	// 	response.Error(r).SetCode(response.DB_READ_ERROR).
	// 		SetMessage(response.CodeMsg(response.DB_READ_ERROR)).Send()
	// }
	// if !garray.NewIntArrayFrom(roleTypes.Ints()).Contains(consts.RoleTypeSuper) {
	// 	response.Error(r).SetCode(response.AUTH_ERROR).
	// 		SetMessage("该职员不是超管，无法登录").Send()
	// }

	// 将当前请求的userID信息保存到请求的上下文c上
	r.SetCtxVar("userId", mc.Id)
	r.Middleware.Next()
}
