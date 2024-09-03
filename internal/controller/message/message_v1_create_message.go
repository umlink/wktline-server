package message

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/message/v1"
)

func (c *ControllerV1) CreateMessage(ctx context.Context, req *v1.CreateMessageReq) (res *v1.CreateMessageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
