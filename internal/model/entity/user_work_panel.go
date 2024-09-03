// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWorkPanel is the golang structure for table user_work_panel.
type UserWorkPanel struct {
	Id         string      `json:"id"         orm:"id"          description:"id"`
	WorkmateId string      `json:"workmateId" orm:"workmate_id" description:"同事 id"`
	UserId     string      `json:"userId"     orm:"user_id"     description:"用户 id(当前用户)"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
