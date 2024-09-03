// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectUser is the golang structure of table project_user for DAO operations like Where/Data.
type ProjectUser struct {
	g.Meta    `orm:"table:project_user, do:true"`
	Id        interface{} //
	ProjectId interface{} // 项目 id
	UserId    interface{} // 用户 id
	Role      interface{} // 用户在项目中的角色
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
