package message

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/message/v1"
)

func (c *ControllerV1) GetMessageList(ctx context.Context, req *v1.GetMessageListReq) (res *v1.GetMessageListRes, err error) {
	out, err := service.Message().GetMessageList(ctx, model.GetMessageListInput{
		Status:    req.Status,
		Type:      req.Type,
		ProjectId: req.ProjectId,
		PageSize:  req.PageSize,
		PageNo:    req.PageNo,
	})
	return (*v1.GetMessageListRes)(out), err
}
