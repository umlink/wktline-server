// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectUser is the golang structure for table project_user.
type ProjectUser struct {
	Id        string      `json:"id"        orm:"id"         description:""`
	ProjectId string      `json:"projectId" orm:"project_id" description:"项目 id"`
	UserId    string      `json:"userId"    orm:"user_id"    description:"用户 id"`
	Role      string      `json:"role"      orm:"role"       description:"用户在项目中的角色"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
