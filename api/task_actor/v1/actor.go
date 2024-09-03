package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskActorItem struct {
	Id       string `json:"id" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
}

type GetTaskActorReq struct {
	g.Meta   `path:"/task/getTaskActor" method:"get" summary:"获取任务下的用户列表" tags:"TaskActor"`
	TaskId   string `json:"taskId" v:"required#项目id不能为空" dc:"项目 id"`
	PageNo   int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int    `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskActorRes struct {
	List     []*TaskActorItem `json:"list" v:"required" dc:"项目用户列表"`
	PageNo   int              `json:"pageNo" v:"required" dc:"页码"`
	PageSize int              `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int              `json:"total" v:"required" dc:"总数"`
}

type AddTaskActorReq struct {
	g.Meta    `path:"/task/addTaskActor" method:"post" summary:"项目添加用户" tags:"TaskActor"`
	ProjectId string   `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	TaskId    string   `json:"taskId" v:"required#任务id不能为空" dc:"任务 id"`
	UserIds   []string `json:"userIds" v:"required"     dc:"用户id列表"`
}
type AddTaskActorRes struct{}

type DelTaskActorReq struct {
	g.Meta    `path:"/task/delTaskActor" method:"post" summary:"删除任务参与人员" tags:"TaskActor"`
	ProjectId string   `json:"projectId" v:"required#项目id不能为空" dc:"项目 id"`
	TaskId    string   `json:"taskId" v:"required#任务id不能为空" dc:"任务 id"`
	UserIds   []string `json:"userIds" v:"required"     dc:"用户ids"`
}
type DelTaskActorRes struct{}
