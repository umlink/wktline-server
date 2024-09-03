package task

import (
	"context"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"

	"wktline-server/api/task/v1"
)

func (c *ControllerV1) UpdateTask(ctx context.Context, req *v1.UpdateTaskReq) (res *v1.UpdateTaskRes, err error) {
	if _, err = service.Task().CheckUpdateForProjectUserAuth(ctx, req.ProjectId); err != nil {
		return nil, err
	}
	err = service.Task().UpdateTask(ctx, model_task.UpdateTaskInput{
		Id:          req.Id,
		ProjectId:   req.ProjectId,
		Name:        req.Name,
		Description: req.Description,
		GroupId:     req.GroupId,
		TypeId:      req.TypeId,
		HandlerId:   req.HandlerId,
		StatusId:    req.StatusId,
		Priority:    req.Priority,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		PlanHour:    req.PlanHour,
		LaborHour:   req.LaborHour,
	})
	return res, err
}
