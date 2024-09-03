package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/user/v1"
)

func (c *ControllerV1) DelUserById(ctx context.Context, req *v1.DelUserByIdReq) (res *v1.DelUserByIdRes, err error) {
	//err = service.User().DeleteUserById(ctx, model.DeleteUserInput{
	//	Id: req.Id,
	//})
	return nil, gerror.New("测试阶段不可删除")
}
