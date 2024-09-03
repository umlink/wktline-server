package middleware

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Log 日志打印
func (s *sMiddleware) Log(r *ghttp.Request) {
	t := time.Now()
	ctx := r.Context()
	r.Middleware.Next()
	// 异步日志
	go func(ctx context.Context, t time.Time) {
		g.Log().Debug(ctx, r.Method, r.RequestURI, time.Since(t))
	}(ctx, t)
}
