package model_task

import (
	v1 "wktline-server/api/task_group/v1"
)

type CreateTaskGroupInput struct {
	Id          string
	ProjectId   string
	GroupName   string
	Description string
	OperatorId  string
}

type DelTaskGroupInput struct {
	Id string
}

type UpdateTaskGroupInput struct {
	Id          string
	GroupName   *string
	Description *string
}

type GetTaskGroupInput struct {
	Keywords  string
	ProjectId string
	PageNo    int
	PageSize  int
}
type GetTaskGroupOutput struct {
	List     []*v1.TaskGroupItem
	PageNo   int
	PageSize int
	Total    int
}
