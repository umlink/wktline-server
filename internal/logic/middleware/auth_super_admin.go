package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"wktline-server/internal/consts"
	"wktline-server/internal/service"
	"wktline-server/utility/response"
)

// AuthSuperAdmin 超管权限
func (s *sMiddleware) AuthSuperAdmin(r *ghttp.Request) {
	role := service.BizCtx().Get(r.Context()).User.Role
	if role != consts.SuperAdmin {
		response.JsonExit(r, 403, "权限不足")
		r.Middleware.Next()
	}
	r.Middleware.Next()
}
