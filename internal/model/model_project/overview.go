package model_project

import v1 "wktline-server/api/project_overview/v1"

type GetProjectStatisticsInput struct {
	ProjectId string
	GroupId   string
	StartTime string
	EndTime   string
}
type GetProjectStatisticsOut struct {
	TaskCount           int
	UserCount           int
	GroupCount          int
	OverTimeDoneCount   int
	OverTimeNoDoneCount int
	LaborHour           float64
	PlanHour            float64
}

type GetProjectUserTaskStatusStatisticsInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectUserTaskStatusStatisticsOutput []v1.ProjectUserTaskStatusStatisticsItem

type GetProjectUserTaskTypeStatisticsInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectUserTaskTypeStatisticsOutput []v1.ProjectUserTaskTypeStatisticsItem

type GetProjectTaskTypeStatInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectTaskTypeStatOutput []v1.ProjectTaskTypeStatItem

type GetProjectTaskStatusStatInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectTaskStatusStatOutput []v1.ProjectTaskStatusStatItem

type GetProjectUserLaborHourStatInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectUserLaborHourStatOutput []v1.ProjectUserLaborHourStatItem

type GetProjectUserTaskCountStatInput struct {
	ProjectId string
	GroupId   string
	UserId    string
	StartTime string
	EndTime   string
}
type GetProjectUserTaskCountStatOutput []v1.ProjectUserTaskCountItem
