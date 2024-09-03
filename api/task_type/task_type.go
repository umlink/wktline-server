// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package task_type

import (
	"context"
	
	"wktline-server/api/task_type/v1"
)

type ITaskTypeV1 interface {
	CreateTaskType(ctx context.Context, req *v1.CreateTaskTypeReq) (res *v1.CreateTaskTypeRes, err error)
	DelTaskType(ctx context.Context, req *v1.DelTaskTypeReq) (res *v1.DelTaskTypeRes, err error)
	UpdateTaskType(ctx context.Context, req *v1.UpdateTaskTypeReq) (res *v1.UpdateTaskTypeRes, err error)
	GetTaskTypeList(ctx context.Context, req *v1.GetTaskTypeListReq) (res *v1.GetTaskTypeListRes, err error)
}


