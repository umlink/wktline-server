package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskGroupItem struct {
	Id          string `json:"id" v:"required" dc:"id"`
	GroupName   string `json:"groupName" v:"required" dc:"名称"`
	Description string `json:"description" v:"required" dc:"描述"`
}

type CreateTaskGroupReq struct {
	g.Meta      `path:"/task/createTaskGroup" method:"post" summary:"创建任务分组" tags:"TaskGroup"`
	ProjectId   string `json:"projectId" v:"required" dc:"项目id"`
	GroupName   string `json:"groupName" v:"required#任务分组名不能为空" dc:"分组名"`
	Description string `json:"description" dc:"描述"`
}
type CreateTaskGroupRes struct {
	Id string `json:"id" v:"required" dc:"分组 id"`
}

type DelTaskGroupReq struct {
	g.Meta    `path:"/task/delTaskGroup" method:"post" summary:"删除任务分组" tags:"TaskGroup"`
	Id        string `json:"id" v:"required#任务分组id不能为空" dc:"任务分组 id id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
}
type DelTaskGroupRes struct{}

type UpdateTaskGroupReq struct {
	g.Meta      `path:"/task/updateTaskGroup" method:"post" summary:"更新任务分组" tags:"TaskGroup"`
	Id          string  `json:"id" v:"required#任务分组id不能为空" dc:"项目分组id"`
	ProjectId   string  `json:"projectId" v:"required" dc:"项目id"`
	GroupName   *string `json:"groupName" dc:"任务分组名"`
	Description *string `json:"description" dc:"任务分组描述"`
}
type UpdateTaskGroupRes struct{}

type GetTaskGroupListReq struct {
	g.Meta    `path:"/task/getTaskGroupList" method:"get" summary:"获取项目分组列表" tags:"TaskGroup"`
	Keywords  string `json:"keywords" dc:"任务分组名称-支持模糊搜索"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskGroupListRes struct {
	List     []*TaskGroupItem `json:"list" v:"required" dc:"任务分组列表"`
	PageNo   int              `json:"pageNo" v:"required" dc:"页码"`
	PageSize int              `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int              `json:"total" v:"required" dc:"总数"`
}
