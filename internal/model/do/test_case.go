// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCase is the golang structure of table test_case for DAO operations like Where/Data.
type TestCase struct {
	g.Meta      `orm:"table:test_case, do:true"`
	Id          interface{} // id
	Name        interface{} // 测试用例名称
	Description interface{} // 说明
	Value       interface{} // 测试用例数据
	Progress    interface{} // 测试进度（1-100）
	ProjectId   interface{} // 项目 id
	CreatorId   interface{} // 创建人
	ActorIds    interface{} // 参与人ids
	Status      interface{} // 状态(草稿: DRAFT 发布:PUBLISHED 测试中:TESTING 测试完成:DONE)
	EditorId    interface{} // 正在编辑的用户 id，lock
	CreatedAt   *gtime.Time //
	DeletedAt   *gtime.Time // 删除时间
}
