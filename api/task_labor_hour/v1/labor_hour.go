package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AddTaskLaborHourReq struct {
	g.Meta      `path:"/task/addLaborHour" method:"post" tags:"TaskLaborHour" summary:"添加任务工时"`
	ProjectId   string    `json:"projectId" v:"required" dc:"项目id"`
	TaskId      string    `json:"taskId" v:"required" dc:"任务id"`
	Date        time.Time `json:"date" v:"date#时间格式错误 YYYY-MM-DD" dc:"日期"`
	Hour        float64   `json:"hour" dc:"计划工时（小时）"`
	Description string    `json:"description" dc:"工作内容说明"`
}
type AddTaskLaborHourRes struct{}

type DelTaskLaborHourReq struct {
	g.Meta    `path:"/task/delLaborHour" method:"post" tags:"TaskLaborHour" summary:"删除任务工时记录"`
	Id        string `json:"id" v:"required" dc:"工时记录的id"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	TaskId    string `json:"taskId" v:"required" dc:"任务id"`
}
type DelTaskLaborHourRes struct{}

type UpdateTaskLaborHourReq struct {
	g.Meta      `path:"/task/updateLaborHour" method:"post" tags:"TaskLaborHour" summary:"更新任务工时记录"`
	ProjectId   string    `json:"projectId" v:"required" dc:"项目id"`
	Id          string    `json:"id" v:"required" dc:"工时记录id"`
	TaskId      string    `json:"taskId" v:"required" dc:"任务id"`
	Date        time.Time `json:"date" v:"date#时间格式错误 YYYY-MM-DD" dc:"日期"`
	Hour        float64   `json:"hour" v:"required" dc:"计划工时（小时）"`
	Description string    `json:"description" v:"required" dc:"工作内容说明"`
}
type UpdateTaskLaborHourRes struct{}

type GetTaskLaborHourReq struct {
	g.Meta   `path:"/task/getLaborHourList" method:"post" tags:"TaskLaborHour" summary:"获取实际工时记录列表"`
	TaskId   string `json:"taskId" v:"required" dc:"任务id"`
	PageNo   int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int    `json:"pageSize" d:"10" dc:"分页数量，最大100"`
}

type TaskLaborHourItem struct {
	Id          string `json:"id" v:"required" dc:"id"`
	UserId      string `json:"userId" dc:"用户id"`
	UserName    string `json:"username" dc:"用户名"`
	Avatar      string `json:"avatar" d:"" dc:"用户头像"`
	Hour        int64  `json:"hour" dc:"时间（小时）"`
	Date        string `json:"date" dc:"日期"`
	Description string `json:"description" dc:"工作内容"`
}
type GetTaskLaborHourRes struct {
	List     []*TaskLaborHourItem `json:"list" v:"required" dc:"工时列表"`
	Total    int                  `json:"total" v:"required" dc:"总数"`
	PageNo   int                  `json:"pageNo" v:"required" dc:"当前页码"`
	PageSize int                  `json:"pageSize" v:"required" dc:"页码大小"`
}
