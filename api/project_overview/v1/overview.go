package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProjectStatisticsReq struct {
	g.Meta    `path:"/project/overview/getStatistics" method:"get" summary:"获取项目下各个指标数量" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	StartTime string `json:"startTime" v:"date#时间格式错误 YYYY-MM-DD" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" v:"date#时间格式错误 YYYY-MM-DD" dc:"任务创建-结束时间"`
}
type GetProjectStatisticsRes struct {
	TaskCount           int     `json:"taskCount" v:"required" dc:"任务数量"`
	UserCount           int     `json:"userCount" v:"required" dc:"任务数量"`
	GroupCount          int     `json:"groupCount" v:"required" dc:"任务分组数"`
	OverTimeDoneCount   int     `json:"overTimeDoneCount" v:"required" dc:"超时完成任务数"`
	OverTimeNoDoneCount int     `json:"overTimeNoDoneCount" v:"required" dc:"超时未完成任务数"`
	LaborHour           float64 `json:"laborHour" v:"required" dc:"实际总工时（h）"`
	PlanHour            float64 `json:"planHour" v:"required" dc:"计划总工时（h）"`
}

type GetProjectUserTaskStatusStatisticsReq struct {
	g.Meta    `path:"/project/overview/getUserTaskStatusStat" method:"get" summary:"获取用户维度的（任务数：按状态）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectUserTaskStatusStatisticsItem struct {
	UserId     string `json:"userId" v:"required" dc:"用户 id"`
	UserName   string `json:"userName" v:"required" dc:"用户名"`
	UserAvatar string `json:"userAvatar" v:"required" dc:"用户头像"`
	StatusId   string `json:"statusId" v:"required" dc:"状态id"`
	StatusName string `json:"statusName" v:"required" dc:"状态名称"`
	TaskCount  int    `json:"taskCount" v:"required" dc:"任务数量"`
}
type GetProjectUserTaskStatusStatisticsRes []ProjectUserTaskStatusStatisticsItem

type GetProjectUserTaskTypeStatisticsReq struct {
	g.Meta    `path:"/project/overview/getUserTaskTypeStat" method:"get" summary:"获取用户维度的（任务数：按类型）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectUserTaskTypeStatisticsItem struct {
	UserId     string `json:"userId" v:"required" dc:"用户 id"`
	UserName   string `json:"userName" v:"required" dc:"用户名"`
	UserAvatar string `json:"userAvatar" v:"required" dc:"用户头像"`
	TypeId     string `json:"typeId" v:"required" dc:"类型id"`
	TypeName   string `json:"typeName" v:"required" dc:"类型名称"`
	TaskCount  int    `json:"taskCount" v:"required" dc:"任务数量"`
}
type GetProjectUserTaskTypeStatisticsRes []ProjectUserTaskTypeStatisticsItem

type GetProjectTaskTypeStatReq struct {
	g.Meta    `path:"/project/overview/getTaskTypeStat" method:"get" summary:"获取项目维度的（任务数：按类型）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectTaskTypeStatItem struct {
	TypeId    string `json:"typeId" v:"required" dc:"类型id"`
	TypeName  string `json:"typeName" v:"required" dc:"类型名称"`
	TaskCount int    `json:"taskCount" v:"required" dc:"任务数量"`
}
type GetProjectTaskTypeStatRes []ProjectTaskTypeStatItem

type GetProjectTaskStatusStatReq struct {
	g.Meta    `path:"/project/overview/getTaskStatusStat" method:"get" summary:"获取项目维度的（任务数：按状态）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectTaskStatusStatItem struct {
	StatusId   string `json:"statusId" v:"required" dc:"状态id"`
	StatusName string `json:"statusName" v:"required" dc:"状态名称"`
	TaskCount  int    `json:"taskCount" v:"required" dc:"任务数量"`
}
type GetProjectTaskStatusStatRes []ProjectTaskStatusStatItem

type GetProjectUserLaborHourReq struct {
	g.Meta    `path:"/project/overview/getUserLaborHourStat" method:"get" summary:"获取项目维度的（各用户工时）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectUserLaborHourStatItem struct {
	UserId     string `json:"userId" v:"required" dc:"用户 id"`
	UserName   string `json:"userName" v:"required" dc:"用户名"`
	UserAvatar string `json:"userAvatar" v:"required" dc:"用户头像"`
	LaborHour  int    `json:"工时" v:"required" dc:"任务工时"`
}
type GetProjectUserLaborHourRes []ProjectUserLaborHourStatItem

type GetProjectUserTaskCountReq struct {
	g.Meta    `path:"/project/overview/getUserTaskCountStat" method:"get" summary:"获取项目维度的（各用户任务数）" tags:"ProjectOverview"`
	ProjectId string `json:"projectId" v:"required" dc:"项目id"`
	GroupId   string `json:"groupId" dc:"分组id"`
	UserId    string `json:"userId" dc:"用户 id"`
	StartTime string `json:"startTime" dc:"任务创建-开始时间"`
	EndTime   string `json:"endTime" dc:"任务创建-结束时间"`
}
type ProjectUserTaskCountItem struct {
	UserId     string `json:"userId" v:"required" dc:"用户 id"`
	UserName   string `json:"userName" v:"required" dc:"用户名"`
	UserAvatar string `json:"userAvatar" v:"required" dc:"用户头像"`
	TaskCount  int    `json:"任务数" v:"required" dc:"任务数量"`
}
type GetProjectUserTaskCountRes []ProjectUserTaskCountItem
