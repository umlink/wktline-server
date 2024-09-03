// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id          string      `json:"id"          orm:"id"           description:""`
	Title       string      `json:"title"       orm:"title"        description:"消息标题"`
	Content     string      `json:"content"     orm:"content"      description:""`
	ContentType string      `json:"contentType" orm:"content_type" description:"内容类型 - （text|markdown|html|扩展）"`
	Status      int         `json:"status"      orm:"status"       description:"消息状态 0：未读，1：已读"`
	MsgType     string      `json:"msgType"     orm:"msg_type"     description:"消息类型"`
	ProjectId   string      `json:"projectId"   orm:"project_id"   description:"项目 id"`
	TaskId      string      `json:"taskId"      orm:"task_id"      description:"任务 id"`
	ReceiverId  string      `json:"receiverId"  orm:"receiver_id"  description:"目标用户 id"`
	SenderId    string      `json:"senderId"    orm:"sender_id"    description:"用户 id"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"是否删除"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
}
