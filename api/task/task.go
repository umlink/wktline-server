// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task

import (
	"context"

	"wktline-server/api/task/v1"
)

type ITaskV1 interface {
	CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error)
	DelTask(ctx context.Context, req *v1.DelTaskReq) (res *v1.DelTaskRes, err error)
	UpdateTask(ctx context.Context, req *v1.UpdateTaskReq) (res *v1.UpdateTaskRes, err error)
	GetTaskDetail(ctx context.Context, req *v1.GetTaskDetailReq) (res *v1.GetTaskDetailRes, err error)
	GetChildTaskList(ctx context.Context, req *v1.GetChildTaskListReq) (res *v1.GetChildTaskListRes, err error)
	GetTaskListByInterval(ctx context.Context, req *v1.GetTaskListByIntervalReq) (res *v1.GetTaskListByIntervalRes, err error)
	GetTaskList(ctx context.Context, req *v1.GetTaskListReq) (res *v1.GetTaskListRes, err error)
}
