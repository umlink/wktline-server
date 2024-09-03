package task_attachment

import (
	"context"

	"wktline-server/utility"

	"wktline-server/api/task_attachment/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) AddTaskAttachment(ctx context.Context, req *v1.AddTaskAttachmentReq) (res *v1.AddTaskAttachmentRes, err error) {
	if _, err := service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	attachmentId := utility.GenUniIdByGuid()
	if err = service.Task().AddTaskAttachment(ctx, model_task.AddTaskAttachmentInput{
		Id:         attachmentId,
		ProjectId:  req.ProjectId,
		TaskId:     req.TaskId,
		ResourceId: req.ResourceId,
		CreatorId:  service.BizCtx().Get(ctx).User.Uid,
	}); err != nil {
		return nil, err
	}
	return &v1.AddTaskAttachmentRes{
		Id: attachmentId,
	}, err
}
