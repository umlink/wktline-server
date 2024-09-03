// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_status

import (
	"context"

	"wktline-server/api/task_status/v1"
)

type ITaskStatusV1 interface {
	CreateTaskStatus(ctx context.Context, req *v1.CreateTaskStatusReq) (res *v1.CreateTaskStatusRes, err error)
	DelTaskStatus(ctx context.Context, req *v1.DelTaskStatusReq) (res *v1.DelTaskStatusRes, err error)
	UpdateTaskStatus(ctx context.Context, req *v1.UpdateTaskStatusReq) (res *v1.UpdateTaskStatusRes, err error)
	UpdateTaskStatusSort(ctx context.Context, req *v1.UpdateTaskStatusSortReq) (res *v1.UpdateTaskStatusSortRes, err error)
	GetTaskStatusList(ctx context.Context, req *v1.GetTaskStatusListReq) (res *v1.GetTaskStatusListRes, err error)
}
