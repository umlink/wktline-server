package middleware

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
	"wktline-server/utility/response"
)

// AuthProject 可编辑项目授权 超管和管理员直接过
func (s *sMiddleware) AuthProject(r *ghttp.Request) {
	role := service.BizCtx().Get(r.Context()).User.Role
	if role == consts.SuperAdmin || role == consts.Admin {
		r.Middleware.Next()
		fmt.Println("系统管理员或系统超管")
	} else {
		// 若存在 projectId 则需检测用户是否为项目负责人或管理员
		projectId := r.GetForm("projectId")
		if !g.IsEmpty(projectId) {
			ctx := r.Context()
			out, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
				ProjectId: gconv.String(projectId),
				UserId:    service.BizCtx().Get(ctx).User.Uid,
			})
			if err != nil {
				response.JsonExit(r, 403, "授权查询异常")
			}
			if out == nil {
				response.JsonExit(r, 403, "未找到对应用户权限")
			}
			if out.Role != consts.User {
				r.Middleware.Next()
			} else {
				response.JsonExit(r, 403, "权限不足")
			}
		}
		r.Middleware.Next()
	}
}
