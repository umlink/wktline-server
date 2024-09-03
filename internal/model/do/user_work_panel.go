// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWorkPanel is the golang structure of table user_work_panel for DAO operations like Where/Data.
type UserWorkPanel struct {
	g.Meta     `orm:"table:user_work_panel, do:true"`
	Id         interface{} // id
	WorkmateId interface{} // 同事 id
	UserId     interface{} // 用户 id(当前用户)
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
