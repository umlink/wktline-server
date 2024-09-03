// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package message

import (
	"context"

	"wktline-server/api/message/v1"
)

type IMessageV1 interface {
	CreateMessage(ctx context.Context, req *v1.CreateMessageReq) (res *v1.CreateMessageRes, err error)
	GetMessageList(ctx context.Context, req *v1.GetMessageListReq) (res *v1.GetMessageListRes, err error)
	ReadMessage(ctx context.Context, req *v1.ReadMessageReq) (res *v1.ReadMessageRes, err error)
	GetMessageDetail(ctx context.Context, req *v1.GetMessageDetailReq) (res *v1.GetMessageDetailRes, err error)
	GetMessageUnReadCount(ctx context.Context, req *v1.GetMessageUnReadCountReq) (res *v1.GetMessageUnReadCountRes, err error)
}
