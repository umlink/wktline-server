// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model"
)

type (
	IMessage interface {
		GetUnreadMessageCount(ctx context.Context) (count int, err error)
		GetMyMessageDetail(ctx context.Context, msgId string) (count int, err error)
		CreateMessage(ctx context.Context, in model.AddMessageInput) (err error)
		CreateTaskMessage(in model.AddMessageInput) (err error)
		CreateProjectActionUsersMsg(ctx context.Context, in model.CreateProjectActionUsersMsgInput) (err error)
		CreateTaskActionUserMsg(ctx context.Context, in model.CreateTaskActionUsersMsgInput) (err error)
		ReadMessage(ctx context.Context, in model.ReadMessageInput) (err error)
		GetMessageList(ctx context.Context, in model.GetMessageListInput) (res *model.GetMessageListOutput, err error)
		GetMessageDetail(ctx context.Context, msgId string) (res *model.GetMessageDetailOutput, err error)
	}
)

var (
	localMessage IMessage
)

func Message() IMessage {
	if localMessage == nil {
		panic("implement not found for interface IMessage, forgot register?")
	}
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}
