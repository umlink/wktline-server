package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogType string

const (
	AddComment             LogType = "COMMENT"                  // 评论
	AddAttachment          LogType = "ATTACHMENT"               // 添加评论附件
	DynamicAddActor        LogType = "DYNAMIC_ADD_ACTOR"        // 添加任务参与者
	DynamicTaskCreate      LogType = "DYNAMIC_TASK_CREATE"      // 修改任务 - 创建
	DynamicTaskName        LogType = "DYNAMIC_TASK_NAME"        // 修改任务 - 名
	DynamicTaskStatus      LogType = "DYNAMIC_TASK_STATUS"      // 修改任务 - 状态
	DynamicTaskType        LogType = "DYNAMIC_TASK_TYPE"        // 修改任务 - 类型
	DynamicTaskPriority    LogType = "DYNAMIC_TASK_PRIORITY"    // 修改任务 - 优先级
	DynamicTaskTime        LogType = "DYNAMIC_TASK_TIME"        // 修改任务 - 时间
	DynamicTaskGroup       LogType = "DYNAMIC_TASK_GROUP"       // 修改任务 - 分组
	DynamicTaskDescription LogType = "DYNAMIC_TASK_DESCRIPTION" // 修改任务 - 描述
	DynamicTaskHandler     LogType = "DYNAMIC_TASK_HANDLER"     // 修改任务 - 处理人
	DynamicTaskLaborTime   LogType = "DYNAMIC_TASK_LABOR_TIME"  // 修改任务 - 实际工时
)

type TaskOperationLogItem struct {
	Id        string  `json:"id" v:"required" dc:"日志id"`
	Name      string  `json:"name" v:"required" dc:"日志名"`
	Content   string  `json:"content" v:"required" dc:"日志内容"`
	Desc      string  `json:"desc" v:"required" dc:"日志备注"`
	Type      LogType `json:"type" v:"required|enums" dc:"日志类型"`
	UserId    string  `json:"userId" v:"required" dc:"用户id"`
	Avatar    string  `json:"avatar" v:"required" dc:"用户头像"`
	Username  string  `json:"username" v:"required" dc:"用户名"`
	CreatedAt string  `json:"createdAt" v:"required" dc:"创建时间"`
}

type CreateTaskOperationLogReq struct {
	g.Meta  `path:"/task/createTaskOperationLog" method:"post" summary:"创建任务类型" tags:"TaskOperationLog"`
	TaskId  string  `json:"taskId" v:"required" dc:"日志id"`
	Name    string  `json:"name" v:"required#日志名不能为空" dc:"日志名/标题"`
	Type    LogType `json:"type" v:"required" dc:"日志名/标题"`
	Content string  `json:"content" dc:"日志内容"`
	Desc    string  `json:"desc" dc:"备注"`
}
type CreateTaskOperationLogRes struct {
	Id int64 `json:"id" v:"required" dc:"任务类型 id"`
}

type DelTaskOperationLogReq struct {
	g.Meta `path:"/task/delTaskOperationLog" method:"post" summary:"删除任务操作日志" tags:"TaskOperationLog"`
	Id     string `json:"id" v:"required#任务类型id不能为空" dc:"任务类型 id id"`
}
type DelTaskOperationLogRes struct{}

type GetTaskOperationLogListReq struct {
	g.Meta   `path:"/task/getTaskOperationLogList" method:"get" summary:"获取任务操作日志" tags:"TaskOperationLog"`
	TaskId   string  `json:"taskId" v:"required#日志id不能为空" dc:"任务id"`
	Type     LogType `json:"type" dc:"日志类型"`
	PageNo   int     `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int     `json:"pageSize" d:"10" dc:"分页数量，最大1000"`
}
type GetTaskOperationLogListRes struct {
	List     []*TaskOperationLogItem `json:"list" v:"required" dc:"任务操作日志列表"`
	PageNo   int                     `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                     `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                     `json:"total" v:"required" dc:"总数"`
}
