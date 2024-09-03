// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_labor_hour

import (
	"context"

	"wktline-server/api/task_labor_hour/v1"
)

type ITaskLaborHourV1 interface {
	AddTaskLaborHour(ctx context.Context, req *v1.AddTaskLaborHourReq) (res *v1.AddTaskLaborHourRes, err error)
	DelTaskLaborHour(ctx context.Context, req *v1.DelTaskLaborHourReq) (res *v1.DelTaskLaborHourRes, err error)
	UpdateTaskLaborHour(ctx context.Context, req *v1.UpdateTaskLaborHourReq) (res *v1.UpdateTaskLaborHourRes, err error)
	GetTaskLaborHour(ctx context.Context, req *v1.GetTaskLaborHourReq) (res *v1.GetTaskLaborHourRes, err error)
}
