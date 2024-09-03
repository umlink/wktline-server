package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskPriorityItem struct {
	Id    string `json:"id" v:"required" dc:"优先级id"`
	Name  string `json:"name" v:"required" dc:"任务优先级"`
	Color string `json:"color" v:"required" dc:"优先级颜色"`
}

type GetTaskPriorityListReq struct {
	g.Meta   `path:"/task/getTaskPriorityList" method:"get" summary:"获取项目优先级列表" tags:"TaskPriority"`
	PageNo   int `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskPriorityListRes struct {
	List     []*TaskPriorityItem `json:"list" v:"required" dc:"任务优先级列表"`
	PageNo   int                 `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                 `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                 `json:"total" v:"required" dc:"总数"`
}

type AddTaskPriorityReq struct {
	g.Meta `path:"/task/addTaskPriority" method:"post" summary:"添加任务优先级" tags:"TaskPriority"`
	Id     string `json:"id" v:"required" dc:"任务优先级"`
	Name   string `json:"name" v:"required" dc:"名称"`
	Color  string `json:"color" v:"required" dc:"优先级颜色"`
}
type AddTaskPriorityRes struct{}

type UpdateTaskPriorityReq struct {
	g.Meta `path:"/task/updateTaskPriority" method:"post" summary:"更新任务优先级" tags:"TaskPriority"`
	Id     string `json:"id" v:"required" dc:"优先级id"`
	Name   string `json:"name" v:"required" dc:"名称"`
	Color  string `json:"color" v:"required" dc:"优先级颜色"`
}
type UpdateTaskPriorityRes struct{}

type DelTaskPriorityReq struct {
	g.Meta `path:"/task/delTaskPriority" method:"post" summary:"删除任务优先级（已占用则不可删除）" tags:"TaskPriority"`
	Id     string `json:"id" v:"required" dc:"优先级id"`
}
type DelTaskPriorityRes struct{}
