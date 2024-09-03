// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_priority

import (
	"context"

	"wktline-server/api/task_priority/v1"
)

type ITaskPriorityV1 interface {
	GetTaskPriorityList(ctx context.Context, req *v1.GetTaskPriorityListReq) (res *v1.GetTaskPriorityListRes, err error)
	AddTaskPriority(ctx context.Context, req *v1.AddTaskPriorityReq) (res *v1.AddTaskPriorityRes, err error)
	UpdateTaskPriority(ctx context.Context, req *v1.UpdateTaskPriorityReq) (res *v1.UpdateTaskPriorityRes, err error)
	DelTaskPriority(ctx context.Context, req *v1.DelTaskPriorityReq) (res *v1.DelTaskPriorityRes, err error)
}
