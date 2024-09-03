// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectInvite is the golang structure for table project_invite.
type ProjectInvite struct {
	Id             string      `json:"id"             orm:"id"               description:"邀请 code"`
	ProjectId      string      `json:"projectId"      orm:"project_id"       description:"项目 id"`
	InviterId      string      `json:"inviterId"      orm:"inviter_id"       description:"发起邀请的用户 id"`
	Deadline       *gtime.Time `json:"deadline"       orm:"deadline"         description:"生效截止时间"`
	MaxInviteCount int         `json:"maxInviteCount" orm:"max_invite_count" description:"z"`
	JoinedCount    int         `json:"joinedCount"    orm:"joined_count"     description:"已加入人数"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:"删除时间"`
}
