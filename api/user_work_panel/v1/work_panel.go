package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddWorkPanelUserReq struct {
	g.Meta  `path:"/work-panel/addUser" method:"post" summary:"工时面板 添加/订阅用户" tags:"UserWorkPanel"`
	UserIds []string `json:"userIds" v:"required" dc:"用户id列表"`
}
type AddWorkPanelUserRes struct{}

type DelWorkPanelUserReq struct {
	g.Meta  `path:"/work-panel/delUser" method:"post" summary:"工时面板 删除用户" tags:"UserWorkPanel"`
	UserIds []string `json:"userIds" v:"required" dc:"用户id列表"`
}
type DelWorkPanelUserRes struct{}

type GetWorkLaborHourLogsReq struct {
	g.Meta    `path:"/work-panel/getLaborHourLogs" method:"post" summary:"获取当前用户添加的用户工时统计" tags:"UserWorkPanel"`
	StartTime string `json:"startTime" v:"datetime#开始时间格式错误 YYYY-MM-DD hh:mm:ss" dc:"任务开始时间"`
	EndTime   string `json:"endTime" v:"datetime|after:StartTime#开始时间格式错误 YYYY-MM-DD HH:mm:ss|开始时间必须小于结束时间" dc:"任务结束时间"`
	PageNo    int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize  int    `json:"pageSize" d:"20" dc:"分页数量，最大100"`
}

type WorkPanelRecordItem struct {
	Hour float64 `json:"hour" v:"required" dc:"工时"`
	Date string  `json:"date" v:"required" dc:"工时日期"`
}

type WorkLaborHourUserItem struct {
	Id       string `json:"id" v:"required" dc:"用户 id"`
	Avatar   string `json:"avatar" v:"required" dc:"用户 头像"`
	UserName string `json:"username" v:"required" dc:"用户名"`
}

type WorkLaborHourLogItem struct {
	User    *WorkLaborHourUserItem `json:"user" v:"required" dc:"用户信息"`
	Records []*WorkPanelRecordItem `json:"records" v:"required" dc:"工作记录"`
}

type GetWorkLaborHourLogsRes struct {
	List     []*WorkLaborHourLogItem `json:"list" v:"required" dc:"数据列表"`
	Total    int                     `json:"total" v:"required" dc:"总数"`
	PageNo   int                     `json:"pageNo" v:"required" dc:"当前页码"`
	PageSize int                     `json:"pageSize" v:"required" dc:"页码大小"`
}

type WorkPanelUserItem struct {
	Id       string `json:"id" v:"required" dc:"用户id"`
	Username string `json:"username" v:"required" dc:"用户名"`
	Nickname string `json:"nickname" v:"required" dc:"用户昵称、别名"`
	Avatar   string `json:"avatar" v:"required" dc:"头像"`
}
type GetWorkPanelUserListReq struct {
	g.Meta   `path:"/work-panel/getUserList" method:"post" summary:"获取用户列表(只获取未订阅的用户)" tags:"UserWorkPanel"`
	Keywords string `json:"keywords" dc:"用户名/用户昵称"`
	PageNo   int    `json:"pageNo" d:"1" dc:"分页号码，默认1"`
	PageSize int    `json:"pageSize" d:"20" dc:"分页数量，最大100"`
}
type GetWorkPanelUserListRes struct {
	List     []*WorkPanelUserItem `json:"list" v:"required" dc:"用户列表"`
	PageNo   int                  `json:"pageNo" v:"required" dc:"页码"`
	PageSize int                  `json:"pageSize" v:"required" dc:"页码大小"`
	Total    int                  `json:"total" v:"required" dc:"总数"`
}
type GetWorkLaborHourByUserIdReq struct {
	g.Meta    `path:"/work-panel/getLaborHourByUserId" method:"post" summary:"获取单个用户的工时记录" tags:"UserWorkPanel"`
	StartTime string `json:"startTime" v:"datetime#开始时间格式错误 YYYY-MM-DD hh:mm:ss" dc:"任务开始时间"`
	EndTime   string `json:"endTime" v:"datetime|after:StartTime#开始时间格式错误 YYYY-MM-DD HH:mm:ss|开始时间必须小于结束时间" dc:"任务结束时间"`
	UserId    string `json:"userId" v:"required"`
}
type WorkLaborHourDetailItem struct {
	Id          string  `json:"id" v:"required" dc:"工时 id"`
	TaskId      string  `json:"taskId" v:"required" dc:"任务 id"`
	TaskName    string  `json:"taskName" v:"required" dc:"任务名"`
	Hour        float64 `json:"hour" v:"required" dc:"工时"`
	Date        string  `json:"date" v:"required" dc:"工时日期"`
	Description string  `json:"description" v:"required" dc:"工作内容"`
}
type GetWorkLaborHourByUserIdRes []*WorkLaborHourDetailItem
