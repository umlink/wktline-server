// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id           string      `json:"id"           orm:"id"            description:"id"`
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"用户头像"`
	Username     string      `json:"username"     orm:"username"      description:"用户名"`
	Nickname     string      `json:"nickname"     orm:"nickname"      description:"用户昵称"`
	Password     string      `json:"password"     orm:"password"      description:"密码"`
	Age          int         `json:"age"          orm:"age"           description:"用户年龄"`
	Six          int         `json:"six"          orm:"six"           description:"用户性别 默认：0， 1：男，2：女，3：未知"`
	Status       int         `json:"status"       orm:"status"        description:"用户状态,默认1 1：正常，0：禁用，2：xxx，3：xxx"`
	Phone        int64       `json:"phone"        orm:"phone"         description:"手机号码"`
	EmpNumber    int64       `json:"empNumber"    orm:"emp_number"    description:"工号"`
	Email        string      `json:"email"        orm:"email"         description:"邮箱"`
	Role         string      `json:"role"         orm:"role"          description:"用户角色 (系统级)USER | ADMIN | SUPER_ADMIN"`
	DepartmentId int64       `json:"departmentId" orm:"department_id" description:"部门 id"`
	IsUpdate     int         `json:"isUpdate"     orm:"is_update"     description:"是否存在更新"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
}
