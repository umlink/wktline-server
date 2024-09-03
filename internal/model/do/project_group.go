// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectGroup is the golang structure of table project_group for DAO operations like Where/Data.
type ProjectGroup struct {
	g.Meta      `orm:"table:project_group, do:true"`
	Id          interface{} // 项目分组 id
	Name        interface{} // 项目分组 id
	Description interface{} // 项目分组描述/简介
	OperatorId  interface{} // 操作者 id
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
