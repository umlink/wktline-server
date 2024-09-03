package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskDetailItem struct {
	Id            string `json:"id" v:"required" dc:"任务id"`
	Name          string `json:"name" v:"required" dc:"任务名称"`
	Description   string `json:"description" v:"required" dc:"任务描述"`
	ProjectId     string `json:"projectId" v:"required" dc:"项目id"`
	ParentId      string `json:"parentId" v:"required" dc:"父任务id"`
	HandlerId     string `json:"handlerId" v:"required" dc:"任务处理者"`
	StatusId      string `json:"statusId" v:"required" dc:"任务状态id"`
	GroupId       string `json:"groupId" v:"required" dc:"分组id"`
	TypeId        string `json:"typeId" v:"required" dc:"任务类型id"`
	Priority      string `json:"priority" v:"required" dc:"优先级"`
	CreatorId     string `json:"creatorId" v:"required" dc:"创建者id"`
	StartTime     string `json:"startTime" v:"required" dc:"任务开始时间"`
	EndTime       string `json:"endTime" v:"required" dc:"任务结束时间"`
	PlanHour      int    `json:"planHour" v:"required" dc:"计划工时(小时)"`
	LaborHour     int    `json:"laborHour" v:"required" dc:"实际使用工时(小时)"`
	SubTaskNum    int    `json:"childrenNum" v:"required" dc:"子任务个数"`
	CreatedAt     string `json:"createdAt" v:"required" dc:"创建时间"`
	UpdatedAt     string `json:"updatedAt" v:"required" dc:"更新时间"`
	StatusName    string `json:"statusName" v:"required" dc:"状态名"`
	StatusEnum    string `json:"statusEnum" v:"required" dc:"状态枚举"`
	StatusColor   string `json:"statusColor" v:"required" dc:"状态颜色"`
	TypeName      string `json:"typeName" v:"required" dc:"任务类型"`
	TypeColor     string `json:"typeColor" v:"required" dc:"任务类型颜色"`
	GroupName     string `json:"groupName" v:"required" dc:"任务分组"`
	CreatorName   string `json:"creatorName" v:"required" dc:"任务创建者名称"`
	CreatorAvatar string `json:"creatorAvatar" v:"required" dc:"任务创建者头像"`
	HandlerName   string `json:"handlerName" v:"required" dc:"任务执行者名称"`
	HandlerAvatar string `json:"handlerAvatar" v:"required" dc:"任务执行者头像"`
}

type CreateTaskReq struct {
	g.Meta      `path:"/task/createTask" method:"post" summary:"新建任务" tags:"Task"`
	Name        *string `json:"name" v:"required" dc:"名称"`
	ProjectId   *string `json:"projectId" v:"required" dc:"项目id"`
	TypeId      *string `json:"typeId" v:"required" dc:"类型id"`
	StatusId    *string `json:"statusId" v:"required" dc:"状态id"`
	Description *string `json:"description" dc:"描述"`
	ParentId    *string `json:"parentId" dc:"父任务id"`
	HandlerId   *string `json:"handlerId" dc:"处理者"`
	GroupId     *string `json:"groupId" dc:"分组id"`
	PriorityId  *string `json:"priorityId" dc:"优先级id"`
	StartTime   string  `json:"startTime" v:"date#开始时间格式不对" dc:"开始时间"`
	EndTime     *string `json:"endTime" v:"date|after-equal:StartTime#结束时间格式不对" dc:"结束时间"`
	PlanHour    *int    `json:"planHour" dc:"计划工时(小时)"`
	LaborHour   *int    `json:"laborHour" dc:"实际使用工时(小时)"`
}

type CreateTaskRes struct {
	Id string `json:"id" v:"required" dc:"任务id"`
}

type DelTaskReq struct {
	g.Meta    `path:"/task/delTask" method:"post" summary:"删除任务" tags:"Task"`
	Id        string `json:"id" v:"required" dc:"任务id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
}
type DelTaskRes struct{}

type UpdateTaskReq struct {
	g.Meta      `path:"/task/updateTask" method:"post" summary:"更新任务" tags:"Task"`
	Id          string  `json:"id" v:"required" dc:"任务id"`
	ProjectId   string  `json:"projectId" v:"required" dc:"项目id"`
	Name        *string `json:"name" dc:"任务名称"`
	Description *string `json:"description" dc:"任务描述"`
	ParentId    *string `json:"parentId" dc:"父任务id"`
	HandlerId   *string `json:"handlerId" dc:"任务处理者"`
	StatusId    *string `json:"statusId" dc:"任务状态id"`
	GroupId     *string `json:"groupId" dc:"分组id"`
	TypeId      *string `json:"typeId" dc:"任务类型id"`
	Priority    *string `json:"priority" dc:"任务优先级id"`
	StartTime   *string `json:"startTime" dc:"任务开始时间"`
	EndTime     *string `json:"endTime" dc:"任务结束时间"`
	PlanHour    *int    `json:"planHour" dc:"计划工时(小时)"`
	LaborHour   *int    `json:"laborHour" dc:"实际使用工时(小时)"`
}
type UpdateTaskRes struct{}

type SortMode struct {
	SortKey string `json:"sortKey" dc:"排序字段"`
	Mode    string `json:"mode" in:"Desc,Asc" dc:"排序模式 Desc|Asc"`
}

type TaskUserItem struct {
	Id       string `json:"id" v:"required" dc:"用户 id"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
	UserName string `json:"userName" v:"required" dc:"用户名"`
}

type GetTaskDetailReq struct {
	g.Meta `path:"/task/getTaskDetail" method:"get" summary:"获取任务详情" tags:"Task"`
	Id     string `json:"id" v:"required" dc:"任务id"`
}

type GetChildTaskListReq struct {
	g.Meta   `path:"/task/getChildTaskList" method:"get" summary:"获取子任务列表" tags:"Task"`
	ParentId string `json:"parentId" v:"required#任务id 不能为空" dc:"任务id"`
	PageNo   int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int    `json:"pageSize" d:"10" dc:"分页数量，最大100"`
}
type GetChildTaskListRes struct {
	List     []*TaskDetailItem `json:"list" v:"required" dc:"任务列表"`
	Total    int               `json:"total" v:"required" dc:"总数"`
	PageNo   int               `json:"pageNo" v:"required" dc:"页码"`
	PageSize int               `json:"pageSize" v:"required" dc:"每页大小"`
}

type GetTaskListByIntervalReq struct {
	g.Meta    `path:"/task/getIntervalTaskList" method:"get" summary:"获取任务列表 - （时间区间)" tags:"Task"`
	StartTime string `json:"startTime" v:"required|datetime" dc:"开始时间"`
	EndTime   string `json:"endTime" v:"required|datetime|after-equal:StartTime" dc:"结束时间"`
}
type TaskListByIntervalItem struct {
	Id          string `json:"id" v:"required" dc:"任务id"`
	Name        string `json:"name" v:"required" dc:"任务名称"`
	ProjectId   string `json:"projectId" v:"required" dc:"项目id"`
	StatusName  string `json:"statusName" v:"required" dc:"状态"`
	StatusEnum  string `json:"statusEnum" v:"required" dc:"状态枚举"`
	StatusColor string `json:"statusColor" v:"required" dc:"状态枚举"`
	StartTime   string `json:"startTime" v:"required" dc:"任务开始时间"`
	EndTime     string `json:"endTime" v:"required" dc:"任务结束时间"`
}
type GetTaskListByIntervalRes []TaskListByIntervalItem

type GetTaskListReq struct {
	g.Meta    `path:"/task/getTaskList" method:"post" summary:"获取任务列表" tags:"Task"`
	Keywords  *string   `json:"keywords" dc:"任务名称(支持模糊查询)"`
	ProjectId *string   `json:"projectId" dc:"项目id"`
	ParentId  *string   `json:"parentId" dc:"父任务id"`
	HandlerId *string   `json:"handlerId" dc:"任务处理者id"`
	CreatorId *string   `json:"creatorId" dc:"任务创建者id"`
	ActorId   *string   `json:"actorId" dc:"参与者用户id"`
	StatusId  *string   `json:"statusId" dc:"任务状态id"`
	GroupId   *string   `json:"groupId" dc:"分组id"`
	TypeId    *string   `json:"typeId" dc:"任务类型id"`
	Priority  *string   `json:"priority" dc:"任务优先级id"`
	SortMode  *SortMode `json:"sortMode" dc:"排序方式"`
	PageNo    int       `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int       `json:"pageSize" d:"10" dc:"分页数量，最大100"`
}

type GetTaskListRes struct {
	List     []*TaskDetailItem `json:"list" v:"required" dc:"任务列表"`
	Total    int               `json:"total" v:"required" dc:"总数"`
	PageNo   int               `json:"pageNo" v:"required" dc:"页码"`
	PageSize int               `json:"pageSize" v:"required" dc:"每页大小"`
}

type GetTaskDetailRes struct {
	TaskDetailItem
}
