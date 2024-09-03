package auth

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

type sJwt struct{}

var authService *jwt.GfJWTMiddleware

func init() {
	jwtKey := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.jwtKey`).String()
	jwtIdentityKey := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.jwtIdentityKey`).String()
	encryptionKey := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.encryptionKey`).String()
	cookieKey := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.cookieKey`).String()
	jwtExpire := g.Cfg().MustGetWithCmd(gctx.New(), `jwtConfig.jwtExpire`).Duration()
	timeout := time.Hour * 24 * jwtExpire
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "JWT_AUTH",
		Key:           []byte(encryptionKey),
		Timeout:       timeout,
		MaxRefresh:    timeout,
		IdentityKey:   "uid",
		TokenLookup:   fmt.Sprintf("header: %s, query: %s, cookie: %s", jwtKey, jwtKey, cookieKey),
		TokenHeadName: jwtIdentityKey,
		TimeFunc:      time.Now,
		Authenticator: Authenticator, // 根据登录信息对用户进行身份验证的回调函数
		// LoginResponse:   LoginResponse,         // 完成登录后返回的信息，用户可自定义返回数据，默认返回
		// RefreshResponse: RefreshResponse,  // 刷新token后返回的信息，用户可自定义返回数据，默认返回
		Unauthorized:    Unauthorized,    // 处理不进行授权的逻辑
		IdentityHandler: IdentityHandler, // 解析并设置用户身份信息
		PayloadFunc:     PayloadFunc,     // 登录期间的回调的函数

	})
	authService = auth
	service.RegisterJwt(New())
}

func New() *sJwt {
	return &sJwt{}
}

func (s *sJwt) Auth() *jwt.GfJWTMiddleware {
	return authService
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator 登录信息写入 token 生成
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	// 含用户 id 时，只为刷新 token
	if user, err := service.User().GetUserByUsernameAndPassword(ctx, in); user != nil {
		return g.Map{
			"uid":      user.Id,
			"username": user.Username,
			"avatar":   user.Avatar,
			"role":     user.Role,
		}, err
	}
	ctxUser := service.BizCtx().Get(ctx).User
	if !g.IsEmpty(ctxUser) {
		userId := ctxUser.Uid
		if !g.IsEmpty(userId) {
			if user, err := service.User().GetUserInfoById(ctx, model.GetUserInfoInput{UserId: userId}); user != nil {
				return g.Map{
					"uid":      user.Id,
					"username": user.Username,
					"avatar":   user.Avatar,
					"role":     user.Role,
				}, err
			}
		}
	}
	return nil, jwt.ErrFailedAuthentication
}
