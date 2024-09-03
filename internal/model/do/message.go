// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta      `orm:"table:message, do:true"`
	Id          interface{} //
	Title       interface{} // 消息标题
	Content     interface{} //
	ContentType interface{} // 内容类型 - （text|markdown|html|扩展）
	Status      interface{} // 消息状态 0：未读，1：已读
	MsgType     interface{} // 消息类型
	ProjectId   interface{} // 项目 id
	TaskId      interface{} // 任务 id
	ReceiverId  interface{} // 目标用户 id
	SenderId    interface{} // 用户 id
	DeletedAt   *gtime.Time // 是否删除
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
