// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task_attachment

import (
	"context"

	"wktline-server/api/task_attachment/v1"
)

type ITaskAttachmentV1 interface {
	AddTaskAttachment(ctx context.Context, req *v1.AddTaskAttachmentReq) (res *v1.AddTaskAttachmentRes, err error)
	DelTaskAttachment(ctx context.Context, req *v1.DelTaskAttachmentReq) (res *v1.DelTaskAttachmentRes, err error)
	GetTaskAttachmentList(ctx context.Context, req *v1.GetTaskAttachmentListReq) (res *v1.GetTaskAttachmentListRes, err error)
}
