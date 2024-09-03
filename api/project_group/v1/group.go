package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProjectGroupItem struct {
	Id          string `json:"id" v:"required" dc:"项目id"`
	Name        string `json:"name" v:"required" dc:"项目名"`
	Description string `json:"description" v:"required" dc:"项目描述"`
}

type CreateProjectGroupReq struct {
	g.Meta      `path:"/project/createProjectGroup" method:"post" summary:"创建项目分组" tags:"ProjectGroup"`
	Name        string `json:"name" v:"required#项目名不能为空" dc:"项目名"`
	Description string `json:"description" dc:"项目描述"`
}
type CreateProjectGroupRes struct {
	Id string `json:"id" v:"required" dc:"项目 id"`
}

type DelProjectGroupReq struct {
	g.Meta `path:"/project/delProjectGroup" method:"post" summary:"删除项目分组" tags:"ProjectGroup"`
	Id     string `json:"id" v:"required#项目分组id不能为空" dc:"项目 id"`
}
type DelProjectGroupRes struct{}

type UpdateProjectGroupReq struct {
	g.Meta      `path:"/project/updateProjectGroup" method:"post" summary:"更新项目分组" tags:"ProjectGroup"`
	Id          string  `json:"id" v:"required" dc:"项目分组id"`
	Name        *string `json:"name" dc:"项目分组名"`
	Description *string `json:"description" dc:"项目分组描述"`
}
type UpdateProjectGroupRes struct{}

type GetProjectGroupListReq struct {
	g.Meta   `path:"/project/getProjectGroupList" method:"get" summary:"获取项目分组列表" tags:"ProjectGroup"`
	Keywords string `json:"keywords" dc:"项目分组名称-支持模糊搜索"`
	PageNo   int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetProjectGroupListRes struct {
	List     []*ProjectGroupItem `json:"list" v:"required" dc:"项目分组列表"`
	PageNo   int                 `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                 `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                 `json:"total" v:"required" dc:"总数"`
}
