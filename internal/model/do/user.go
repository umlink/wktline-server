// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	Id           interface{} // id
	Avatar       interface{} // 用户头像
	Username     interface{} // 用户名
	Nickname     interface{} // 用户昵称
	Password     interface{} // 密码
	Age          interface{} // 用户年龄
	Six          interface{} // 用户性别 默认：0， 1：男，2：女，3：未知
	Status       interface{} // 用户状态,默认1 1：正常，0：禁用，2：xxx，3：xxx
	Phone        interface{} // 手机号码
	EmpNumber    interface{} // 工号
	Email        interface{} // 邮箱
	Role         interface{} // 用户角色 (系统级)USER | ADMIN | SUPER_ADMIN
	DepartmentId interface{} // 部门 id
	IsUpdate     interface{} // 是否存在更新
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
