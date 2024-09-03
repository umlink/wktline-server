package message

import (
	"context"

	"wktline-server/api/message/v1"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetMessageUnReadCount(ctx context.Context, req *v1.GetMessageUnReadCountReq) (res *v1.GetMessageUnReadCountRes, err error) {
	count, err := service.Message().GetUnreadMessageCount(ctx)
	return &v1.GetMessageUnReadCountRes{
		Count: count,
	}, err
}
