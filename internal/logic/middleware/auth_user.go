package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"wktline-server/internal/service"
	"wktline-server/utility/response"
)

func (s *sMiddleware) AuthUser(r *ghttp.Request) {
	if g.IsEmpty(service.BizCtx().Get(r.Context())) {
		response.JsonExit(r, 401, "请先登录!")
	}
	r.Middleware.Next()
}
