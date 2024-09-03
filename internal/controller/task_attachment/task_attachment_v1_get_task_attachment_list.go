package task_attachment

import (
	"context"

	"wktline-server/api/task_attachment/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) GetTaskAttachmentList(ctx context.Context, req *v1.GetTaskAttachmentListReq) (res *v1.GetTaskAttachmentListRes, err error) {
	out, err := service.Task().GetTaskAttachmentList(ctx, model_task.GetTaskAttachmentListInput{
		TaskId:    req.TaskId,
		ProjectId: req.ProjectId,
		CreatorId: req.CreatorId,
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
	})
	return &v1.GetTaskAttachmentListRes{
		List:     out.List,
		Total:    out.Total,
		PageSize: out.PageSize,
		PageNo:   out.PageNo,
	}, err
}
