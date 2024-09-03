package model_task

import (
	v1 "wktline-server/api/task/v1"
)

type CreateTaskInput struct {
	Id          string
	Name        *string
	ProjectId   *string
	StatusId    *string
	TypeId      *string
	CreatorId   string
	Description *string
	HandlerId   *string
	ParentId    *string
	GroupId     *string
	PriorityId  *string
	StartTime   string
	EndTime     *string
	PlanHour    *int
	LaborHour   *int
}
type CreateTaskOutput string

type DelTaskInput struct {
	Id string
}

type UpdateTaskInput struct {
	Id           string
	ProjectId    string
	Name         *string
	Description  *string
	ParentId     *string
	HandlerId    *string
	StatusId     *string
	GroupId      *string
	TypeId       *string
	Priority     *string
	StartTime    *string
	FinishedTime string
	EndTime      *string
	PlanHour     *int
	LaborHour    *int
}

type GetTaskDetailInput struct {
	Id string
}
type GetTaskDetailOutput v1.GetTaskDetailRes
type GetTaskListInput struct {
	ProjectId *string
	Keywords  *string
	ParentId  *string
	HandlerId *string
	StatusId  *string
	GroupId   *string
	TypeId    *string
	Priority  *string
	SortMode  *v1.SortMode `json:"sortMode" dc:"排序方式"`
	PageNo    int
	PageSize  int
}
type GetTaskListOutput v1.GetTaskListRes

type GetChildTaskListInput struct {
	ParentId string
	PageNo   int
	PageSize int
}
type GetChildTaskListOutput v1.GetTaskListRes

type GetTaskListByIntervalInput struct {
	StartTime string
	EndTime   string
}
type GetTaskListByIntervalOutput []v1.TaskListByIntervalItem
