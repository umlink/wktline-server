// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_actor

import (
	"context"

	"wktline-server/api/task_actor/v1"
)

type ITaskActorV1 interface {
	GetTaskActor(ctx context.Context, req *v1.GetTaskActorReq) (res *v1.GetTaskActorRes, err error)
	AddTaskActor(ctx context.Context, req *v1.AddTaskActorReq) (res *v1.AddTaskActorRes, err error)
	DelTaskActor(ctx context.Context, req *v1.DelTaskActorReq) (res *v1.DelTaskActorRes, err error)
}
