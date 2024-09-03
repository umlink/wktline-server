// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package task_group

import (
	"context"
	
	"wktline-server/api/task_group/v1"
)

type ITaskGroupV1 interface {
	CreateTaskGroup(ctx context.Context, req *v1.CreateTaskGroupReq) (res *v1.CreateTaskGroupRes, err error)
	DelTaskGroup(ctx context.Context, req *v1.DelTaskGroupReq) (res *v1.DelTaskGroupRes, err error)
	UpdateTaskGroup(ctx context.Context, req *v1.UpdateTaskGroupReq) (res *v1.UpdateTaskGroupRes, err error)
	GetTaskGroupList(ctx context.Context, req *v1.GetTaskGroupListReq) (res *v1.GetTaskGroupListRes, err error)
}


