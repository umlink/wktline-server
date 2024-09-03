// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package task_operation_log

import (
	"context"
	
	"wktline-server/api/task_operation_log/v1"
)

type ITaskOperationLogV1 interface {
	CreateTaskOperationLog(ctx context.Context, req *v1.CreateTaskOperationLogReq) (res *v1.CreateTaskOperationLogRes, err error)
	DelTaskOperationLog(ctx context.Context, req *v1.DelTaskOperationLogReq) (res *v1.DelTaskOperationLogRes, err error)
	GetTaskOperationLogList(ctx context.Context, req *v1.GetTaskOperationLogListReq) (res *v1.GetTaskOperationLogListRes, err error)
}


