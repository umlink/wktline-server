package task

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"wktline-server/internal/dao"
	"wktline-server/internal/model/model_task"
)

func (s *sTask) CreateTaskGroup(ctx context.Context, in model_task.CreateTaskGroupInput) (err error) {
	_, err = dao.TaskGroup.Ctx(ctx).Data(in).OmitEmptyData().Insert()
	return err
}

func (s *sTask) DelTaskGroup(ctx context.Context, in model_task.DelTaskGroupInput) (err error) {
	_, err = dao.TaskGroup.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sTask) UpdateTaskGroup(ctx context.Context, in model_task.UpdateTaskGroupInput) (err error) {
	_, err = dao.TaskGroup.Ctx(ctx).Data(in).WherePri(in.Id).Update()
	return err
}

func (s *sTask) GetTaskGroup(ctx context.Context, in model_task.GetTaskGroupInput) (res *model_task.GetTaskGroupOutput, err error) {
	out := model_task.GetTaskGroupOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.TaskGroup.Ctx(ctx).Where("project_id", in.ProjectId)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.WhereLike("group_name", likePattern)
	}
	if err = d.Page(in.PageNo, in.PageSize).
		Fields("id,group_name,description").
		OrderDesc("created_at").ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	return &out, err
}
