// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectInvite is the golang structure of table project_invite for DAO operations like Where/Data.
type ProjectInvite struct {
	g.Meta         `orm:"table:project_invite, do:true"`
	Id             interface{} // 邀请 code
	ProjectId      interface{} // 项目 id
	InviterId      interface{} // 发起邀请的用户 id
	Deadline       *gtime.Time // 生效截止时间
	MaxInviteCount interface{} // z
	JoinedCount    interface{} // 已加入人数
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
