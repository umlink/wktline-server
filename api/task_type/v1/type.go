package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskTypeItem struct {
	Id    string `json:"id" v:"required" dc:"类型id"`
	Name  string `json:"name" v:"required" dc:"任务类型"`
	Color string `json:"color" v:"required" dc:"类型颜色"`
}

type CreateTaskTypeReq struct {
	g.Meta    `path:"/task/createTaskType" method:"post" summary:"创建任务类型" tags:"TaskType"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	Name      string `json:"name" v:"required#任务类型名不能为空" dc:"类型"`
	Color     string `json:"color" v:"required#颜色不能为空" dc:"类型颜色"`
}
type CreateTaskTypeRes struct {
	Id string `json:"id" v:"required" dc:"任务类型 id"`
}

type DelTaskTypeReq struct {
	g.Meta    `path:"/task/delTaskType" method:"post" summary:"删除任务类型" tags:"TaskType"`
	Id        string `json:"id" v:"required#任务类型id不能为空" dc:"任务类型 id id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
}
type DelTaskTypeRes struct{}

type UpdateTaskTypeReq struct {
	g.Meta    `path:"/task/updateTaskType" method:"post" summary:"更新任务类型" tags:"TaskType"`
	Id        string `json:"id" v:"required#任务类型id不能为空" dc:"项目类型id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	Name      string `json:"name" dc:"名称"`
	Color     string `json:"color" dc:"颜色"`
}
type UpdateTaskTypeRes struct{}

type GetTaskTypeListReq struct {
	g.Meta    `path:"/task/getTaskTypeList" method:"get" summary:"获取项目类型列表" tags:"TaskType"`
	Keywords  string `json:"keywords" dc:"任务类型名称-支持模糊搜索"`
	ProjectId string `json:"projectId" v:"required#项目id不能为空" dc:"项目id"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskTypeListRes struct {
	List     []*TaskTypeItem `json:"list" v:"required" dc:"任务类型列表"`
	PageNo   int             `json:"pageNo" v:"required" dc:"页码"`
	PageSize int             `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int             `json:"total" v:"required" dc:"总数"`
}
