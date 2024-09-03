package task

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
)

func (s *sTask) AddTaskAttachment(ctx context.Context, in model_task.AddTaskAttachmentInput) (err error) {
	_, err = dao.TaskAttachment.Ctx(ctx).Data(in).OmitEmptyData().Insert()
	return err
}

func (s *sTask) DelTaskAttachment(ctx context.Context, in model_task.DelTaskAttachmentInput) (err error) {
	_, err = dao.TaskAttachment.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sTask) GetTaskAttachmentList(ctx context.Context, in model_task.GetTaskAttachmentListInput) (res *model_task.GetTaskAttachmentListOutput, err error) {
	out := model_task.GetTaskAttachmentListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := g.Model("resource r,task_attachment ta").Where("ta.resource_id=r.id")
	if !g.IsEmpty(in.ProjectId) {
		d = d.Where("ta.project_id", in.ProjectId)
	}
	if !g.IsEmpty(in.TaskId) {
		d = d.Where("ta.task_id", in.TaskId)
	}
	if !g.IsEmpty(in.CreatorId) {
		d = d.Where("ta.creator_id", in.CreatorId)
	}
	err = d.Page(in.PageNo, in.PageSize).Fields("ta.id,r.name,r.url,r.type,r.size,r.hash").ScanAndCount(&out.List, &out.Total, false)
	if err != nil {
		return nil, err
	}
	return &out, err
}
