package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/utility"

	v1 "wktline-server/api/login/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (user *v1.LoginRes, err error) {
	out, err := service.User().GetUserByUsernameAndPassword(ctx, model.UserLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if g.IsEmpty(out) || err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
	}
	out.Token, out.Expire = service.Jwt().Auth().LoginHandler(ctx)
	utility.SetCookie(ctx, out.Token)
	return out, nil
}
