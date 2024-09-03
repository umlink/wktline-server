package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"wktline-server/internal/service"
)

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Jwt().Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
