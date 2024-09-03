package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	"net/http"
	"time"
)

// GenUniIdByGuid 生成唯一 id
func GenUniIdByGuid() string {
	return guid.S([]byte("ZhaoFuLin"))
}

func SetCookie(ctx context.Context, token string) {
	r := g.RequestFromCtx(ctx)
	cookieKey := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.cookieKey`).String()
	jwtExpire := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.jwtExpire`).Duration()
	r.Cookie.SetHttpCookie(&http.Cookie{
		Name:     cookieKey,
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour * jwtExpire),
		Path:     "/",
		HttpOnly: false,
	})
}
