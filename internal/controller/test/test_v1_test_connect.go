package test

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"wktline-server/api/test/v1"
)

func (c *ControllerV1) TestConnect(ctx context.Context, req *v1.TestConnectReq) (res *v1.TestConnectRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
