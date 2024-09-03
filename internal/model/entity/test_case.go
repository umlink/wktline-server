// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCase is the golang structure for table test_case.
type TestCase struct {
	Id          string      `json:"id"          orm:"id"          description:"id"`
	Name        string      `json:"name"        orm:"name"        description:"测试用例名称"`
	Description string      `json:"description" orm:"description" description:"说明"`
	Value       string      `json:"value"       orm:"value"       description:"测试用例数据"`
	Progress    int         `json:"progress"    orm:"progress"    description:"测试进度（1-100）"`
	ProjectId   string      `json:"projectId"   orm:"project_id"  description:"项目 id"`
	CreatorId   string      `json:"creatorId"   orm:"creator_id"  description:"创建人"`
	ActorIds    string      `json:"actorIds"    orm:"actor_ids"   description:"参与人ids"`
	Status      string      `json:"status"      orm:"status"      description:"状态(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)"`
	EditorId    string      `json:"editorId"    orm:"editor_id"   description:"正在编辑的用户 id，lock"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
