package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateMessageReq struct {
	g.Meta     `path:"/message/addMsg" method:"post" summary:"新增一条消息" tags:"Message"`
	Title      string `json:"title" v:"required" dc:"消息标题"`
	Content    string `json:"content" v:"required" dc:"消息内容"`
	Type       string `json:"type" v:"required" dc:"消息类型"`
	ProjectId  string `json:"projectId" dc:"项目id"`
	TaskId     string `json:"taskId" dc:"任务id"`
	ReceiverId string `json:"receiverId" v:"required" dc:"接收者 id"`
	SenderId   string `json:"senderId" v:"required" dc:"发送者 id"`
}
type CreateMessageRes struct{}

type GetMessageListReq struct {
	g.Meta    `path:"/message/getMsgList" method:"get" summary:"获取消息列表" tags:"Message"`
	Type      string `json:"type" dc:"消息类型"`
	Status    string `json:"status" dc:"已读未读 1｜0"`
	ProjectId string `json:"projectId" dc:"项目id"`
	PageNo    int    `json:"pageNo" d:"1" dc:"页码"`
	PageSize  int    `json:"pageSize" d:"10" dc:"页码大小"`
}
type MessageItem struct {
	Id           string `json:"id" v:"required" dc:"消息id"`
	Title        string `json:"title" v:"required" dc:"消息标题"`
	Content      string `json:"content" v:"required" dc:"内容"`
	TaskId       string `json:"taskId" v:"required" dc:"任务 id"`
	TaskName     string `json:"taskName" v:"required" dc:"任务标题"`
	Status       int    `json:"status" v:"required" dc:"已读未读 1｜0"`
	MsgType      string `json:"msgType" v:"required" dc:"消息类型"`
	SenderId     string `json:"senderId" v:"required" dc:"发送者 id"`
	SenderName   string `json:"senderName" v:"required" dc:"发送者名"`
	SenderAvatar string `json:"senderAvatar" v:"required" dc:"发送者头像"`
	CreatedAt    string `json:"createdAt" v:"required" dc:"消息创建时间"`
}
type GetMessageListRes struct {
	List     []*MessageItem `json:"list" v:"required" dc:"消息列表"`
	PageNo   int            `json:"pageNo" v:"required" dc:"页码"`
	PageSize int            `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int            `json:"total" v:"required" dc:"总数"`
}

type ReadMessageReq struct {
	g.Meta `path:"/message/readMsg" method:"post" summary:"已读消息" tags:"Message"`
	Id     string `json:"id" v:"required" dc:"消息id"`
}
type ReadMessageRes struct{}

type GetMessageDetailReq struct {
	g.Meta `path:"/message/getMsgDetail" method:"post" summary:"消息详情" tags:"Message"`
	Id     string `json:"id" v:"required" dc:"消息id"`
}
type GetMessageDetailRes struct {
	Id             string `json:"id" v:"required" dc:"消息id"`
	Title          string `json:"title" v:"required" dc:"消息标题"`
	Content        string `json:"content" v:"required" dc:"消息内容"`
	TaskId         string `json:"taskId" dc:"任务id"`
	TaskName       string `json:"taskName" v:"required" dc:"任务标题"`
	ContentType    string `json:"contentType" v:"required" dc:"内容格式类型"`
	MsgType        string `json:"msgType" v:"required" dc:"消息类型"`
	ProjectId      string `json:"projectId" dc:"项目id"`
	ReceiverId     string `json:"receiverId" v:"required" dc:"接收者 id"`
	ReceiverName   string `json:"receiverName" v:"required" dc:"接收者名"`
	ReceiverAvatar string `json:"receiverAvatar" v:"required" dc:"接收者头像"`
	SenderId       string `json:"senderId" v:"required" dc:"发送者 id"`
	SenderName     string `json:"senderName" v:"required" dc:"发送者名"`
	SenderAvatar   string `json:"senderAvatar" v:"required" dc:"发送者头像"`
	CreatedAt      string `json:"createdAt" v:"required" dc:"消息创建时间"`
}
type GetMessageUnReadCountReq struct {
	g.Meta `path:"/message/getMsgUnreadCount" method:"post" summary:"消息详情" tags:"Message"`
}
type GetMessageUnReadCountRes struct {
	Count int `json:"count" v:"required" dc:"消息未读数"`
}
