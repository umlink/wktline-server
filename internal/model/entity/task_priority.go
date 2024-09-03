// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskPriority is the golang structure for table task_priority.
type TaskPriority struct {
	Id        string      `json:"id"        orm:"id"         description:"优先级"`
	Name      string      `json:"name"      orm:"name"       description:"优先级名"`
	Color     string      `json:"color"     orm:"color"      description:"颜色"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
