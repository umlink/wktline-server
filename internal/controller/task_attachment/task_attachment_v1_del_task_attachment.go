package task_attachment

import (
	"context"

	"wktline-server/api/task_attachment/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTaskAttachment(ctx context.Context, req *v1.DelTaskAttachmentReq) (res *v1.DelTaskAttachmentRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	err = service.Task().DelTaskAttachment(ctx, model_task.DelTaskAttachmentInput{
		Id: req.Id,
	})
	return nil, err
}
