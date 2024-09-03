package middleware

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"

	"wktline-server/internal/consts"
	"wktline-server/internal/service"
	"wktline-server/utility/response"
)

// AuthAdmin 管理员权限
func (s *sMiddleware) AuthAdmin(r *ghttp.Request) {
	role := service.BizCtx().Get(r.Context()).User.Role
	fmt.Println("用户权限：", role)
	if role == consts.User {
		response.JsonExit(r, 403, "权限不足")
	}
	r.Middleware.Next()
}
