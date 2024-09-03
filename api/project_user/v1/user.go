package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProjectUserItem struct {
	Id       string `json:"id" v:"required" dc:"id"`
	UserId   string `json:"userId" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Email    string `json:"email" v:"required" dc:"邮箱"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
	Role     string `json:"role" v:"required" dc:"用户角色"`
}

type GetProjectUserReq struct {
	g.Meta    `path:"/project/getProjectUser" method:"get" summary:"获取项目下的用户列表" tags:"ProjectUser"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	Keywords  string `json:"keywords" dc:"用户名，支持模糊搜索"`
	Role      string `json:"role" dc:"用户权限类型"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetProjectUserRes struct {
	List     []*ProjectUserItem `json:"list" v:"required" dc:"项目用户列表"`
	PageNo   int                `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                `json:"total" v:"required" dc:"总数"`
}

type AddProjectUserByIdsReq struct {
	g.Meta    `path:"/project/addProjectUserByIds" method:"post" summary:"项目项目人员" tags:"ProjectUser"`
	ProjectId string   `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	UserIds   []string `json:"userIds" v:"required" dc:"用户id列表"`
}
type AddProjectUserByIdsRes struct{}

type DelProjectUserByIdsReq struct {
	g.Meta    `path:"/project/delProjectUserByIds" method:"post" summary:"删除项目内人员" tags:"ProjectUser"`
	ProjectId string   `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	UserIds   []string `json:"userIds"  v:"required" dc:"用户id列表"`
}
type DelProjectUserByIdsRes struct{}

type UpdateProjectUserRoleReq struct {
	g.Meta    `path:"/project/updateProjectUserRole" method:"post" summary:"更新项目成员角色" tags:"ProjectUser"`
	UserId    string `json:"userId" v:"required#用户id不能为空" dc:"用户 id"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	Role      string `json:"role" v:"required|in:USER,ADMIN#用户角色不能为空｜role可写范围为USER,ADMIN" dc:"用户角色"`
}
type UpdateProjectUserRoleRes struct{}
