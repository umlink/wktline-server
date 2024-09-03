// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskAttachment is the golang structure of table task_attachment for DAO operations like Where/Data.
type TaskAttachment struct {
	g.Meta     `orm:"table:task_attachment, do:true"`
	Id         interface{} // id
	ProjectId  interface{} // 项目 id
	TaskId     interface{} // 任务 id
	ResourceId interface{} // 资源 id
	CreatorId  interface{} // 创建者 id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
