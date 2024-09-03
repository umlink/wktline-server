package message

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/message/v1"
)

func (c *ControllerV1) ReadMessage(ctx context.Context, req *v1.ReadMessageReq) (res *v1.ReadMessageRes, err error) {
	count, err := service.Message().GetMyMessageDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("您没有这条消息")
	}
	return nil, service.Message().ReadMessage(ctx, model.ReadMessageInput{
		Id: req.Id,
	})
}
