package model_task

import (
	v1 "wktline-server/api/task_type/v1"
)

type CreateTaskTypeInput struct {
	Id         string
	Name       string
	Color      string
	ProjectId  string
	OperatorId string
}

type InitTaskTypeInput struct {
	Id         string
	Name       string
	Color      string
	ProjectId  string
	OperatorId string
}

type DelTaskTypeInput struct {
	Id        string
	ProjectId string
}

type UpdateTaskTypeInput struct {
	ProjectId string
	Id        string
	Name      string
	Color     string
}

type GetTaskTypeInput struct {
	Keywords  string
	ProjectId string
	PageNo    int
	PageSize  int
}
type GetTaskTypeOutput struct {
	List     []*v1.TaskTypeItem
	PageNo   int
	PageSize int
	Total    int
}
