package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"

	"wktline-server/api/task/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error) {
	user := service.BizCtx().Get(ctx).User
	pUser, _ := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
		ProjectId: *req.ProjectId,
		UserId:    user.Uid,
	})
	// 普通用户必须在项目中才可新建任务
	if user.Role == consts.User && g.IsEmpty(pUser) {
		return nil, gerror.New("非本项目成员无权新建任务")
	}
	taskId, err := service.Task().CreateTask(ctx, model_task.CreateTaskInput{
		Name:        req.Name,
		Description: req.Description,
		ParentId:    req.ParentId,
		ProjectId:   req.ProjectId,
		HandlerId:   req.HandlerId,
		StatusId:    req.StatusId,
		GroupId:     req.GroupId,
		TypeId:      req.TypeId,
		PriorityId:  req.PriorityId,
		PlanHour:    req.PlanHour,
		LaborHour:   req.LaborHour,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		CreatorId:   service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.CreateTaskRes{
		Id: taskId,
	}, err
}
