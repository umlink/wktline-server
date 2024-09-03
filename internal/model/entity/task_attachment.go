// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskAttachment is the golang structure for table task_attachment.
type TaskAttachment struct {
	Id         string      `json:"id"         orm:"id"          description:"id"`
	ProjectId  string      `json:"projectId"  orm:"project_id"  description:"项目 id"`
	TaskId     string      `json:"taskId"     orm:"task_id"     description:"任务 id"`
	ResourceId string      `json:"resourceId" orm:"resource_id" description:"资源 id"`
	CreatorId  string      `json:"creatorId"  orm:"creator_id"  description:"创建者 id"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
}
