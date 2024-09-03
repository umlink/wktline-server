package message

import (
	"context"
	"wktline-server/internal/service"

	"wktline-server/api/message/v1"
)

func (c *ControllerV1) GetMessageDetail(ctx context.Context, req *v1.GetMessageDetailReq) (res *v1.GetMessageDetailRes, err error) {
	out, err := service.Message().GetMessageDetail(ctx, req.Id)
	return (*v1.GetMessageDetailRes)(out), err
}
