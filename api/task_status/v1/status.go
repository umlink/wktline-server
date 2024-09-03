package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskStatusItem struct {
	Id      string `json:"id" v:"required" dc:"状态id"`
	Name    string `json:"name" v:"required" dc:"任务状态"`
	Color   string `json:"color" v:"required" dc:"状态颜色"`
	Sort    int    `json:"sort" v:"required" dc:"排序"`
	Enum    string `json:"enum" v:"required" dc:"枚举"`
	Default int    `json:"default" v:"required" dc:"排序"`
}

type CreateTaskStatusReq struct {
	g.Meta    `path:"/task/createTaskStatus" method:"post" summary:"创建任务状态" tags:"TaskStatus"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	Name      string `json:"name" v:"required#任务状态名不能为空" dc:"状态"`
	Color     string `json:"color" v:"required#颜色不能为空" dc:"状态颜色"`
	Enum      string `json:"enum" v:"required#枚举不能为空" dc:"枚举"`
}
type CreateTaskStatusRes struct {
	Id string `json:"id" v:"required" dc:"任务状态 id"`
}

type DelTaskStatusReq struct {
	g.Meta    `path:"/task/delTaskStatus" method:"post" summary:"删除任务状态" tags:"TaskStatus"`
	Id        string `json:"id" v:"required#任务状态id不能为空" dc:"任务状态 id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目 id"`
}
type DelTaskStatusRes struct{}

type UpdateTaskStatusReq struct {
	g.Meta    `path:"/task/updateTaskStatus" method:"post" summary:"更新任务状态详情" tags:"TaskStatus"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	Id        string `json:"id" v:"required#任务状态id不能为空" dc:"任务状态id"`
	Name      string `json:"name" dc:"名称"`
	Color     string `json:"color" dc:"颜色"`
	Sort      int    `json:"sort" v:"required" dc:"排序"`
	Enum      string `json:"enum" dc:"枚举值"`
}
type UpdateTaskStatusRes struct{}

type SortMapItem struct {
	Id   string `json:"id" v:"required#任务状态id不能为空" dc:"项目状态id"`
	Sort int    `json:"sort" v:"required" dc:"排序"`
}
type UpdateTaskStatusSortReq struct {
	g.Meta      `path:"/task/updateTaskStatusSort" method:"post" summary:"更新任务状态排序" tags:"TaskStatus"`
	SortMapList []*SortMapItem `json:"sortMapList" dc:"项目id"`
	ProjectId   string         `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
}
type UpdateTaskStatusSortRes struct{}

type GetTaskStatusListReq struct {
	g.Meta    `path:"/task/getTaskStatusList" method:"get" summary:"获取项目状态列表" tags:"TaskStatus"`
	Keywords  string `json:"keywords" dc:"任务状态名称-支持模糊搜索"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskStatusListRes struct {
	List     []*TaskStatusItem `json:"list" v:"required" dc:"任务状态列表"`
	PageNo   int               `json:"pageNo" v:"required" dc:"页码"`
	PageSize int               `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int               `json:"total" v:"required" dc:"总数"`
}
