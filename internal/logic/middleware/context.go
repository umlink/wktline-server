package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"wktline-server/internal/consts"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

// Context 自定义上下文对象
func (s *sMiddleware) Context(r *ghttp.Request) {
	token := r.Cookie.Get(consts.UserCookieKey)
	cache, err := g.Redis().Get(r.Context(), gconv.String(token))
	if !g.IsEmpty(cache) && err == nil {
		user := &model.ContextPayloadUser{}
		err = gconv.Struct(cache, &user)
		// 设置上下文数据
		customCtx := &model.Context{
			Data: make(g.Map),
			User: user,
		}
		service.BizCtx().Init(r, customCtx)
	}
	r.Middleware.Next()
}
