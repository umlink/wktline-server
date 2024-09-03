package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility/response"
)

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	var user = &model.ContextPayloadUser{}
	err := gconv.Scan(gconv.Map(service.Jwt().Auth().GetPayload(r.Context())), user)
	if err != nil {
		response.JsonExit(r, 401, "拒绝访问")
		r.Middleware.Next()
	}
	customCtx := &model.Context{
		User: &model.ContextPayloadUser{
			Uid:      user.Uid,
			Username: user.Username,
			Avatar:   user.Avatar,
			Role:     user.Role,
		},
	}
	service.BizCtx().Init(r, customCtx)
	r.Middleware.Next()
}
