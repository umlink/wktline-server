// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Project is the golang structure of table project for DAO operations like Where/Data.
type Project struct {
	g.Meta      `orm:"table:project, do:true"`
	Id          interface{} // 项目 id
	Name        interface{} // 项目名
	Description interface{} // 项目描述
	HeaderImg   interface{} // 项目头图
	CreatorId   interface{} // 创建人 id
	OwnerId     interface{} // 项目所有者 id
	OperatorId  interface{} // 项目操作者 id
	GroupId     interface{} // 项目分组 id
	Status      interface{} // 项目状态 1: 正常 2: 删除
	ShowType    interface{} // 显示类型 PUBLIC: 公开 PRIVATE: 私有
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time // 删除shi
}
