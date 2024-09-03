package model_task

import (
	v1 "wktline-server/api/task_status/v1"
)

type CreateTaskStatusInput struct {
	Id         string
	Sort       int
	Name       string
	Color      string
	Enum       string
	ProjectId  string
	OperatorId string
}

type InitTaskStatusInput struct {
	Id         string
	Sort       int
	Name       string
	Color      string
	Enum       string
	ProjectId  string
	Default    int
	OperatorId string
}

type DelTaskStatusInput struct {
	Id        string
	ProjectId string
}

type UpdateTaskStatusInput struct {
	Id        string
	Name      string
	Color     string
	Enum      string
	ProjectId string
}

type GetTaskStatusByIdOutput struct {
	Id    string
	Name  string
	Color string
	Enum  string
}

type UpdateTaskStatusSortInput struct {
	ProjectId   string
	SortMapList []*v1.SortMapItem
}

type GetTaskStatusInput struct {
	ProjectId string
	PageNo    int
	PageSize  int
}
type GetTaskStatusOutput struct {
	List     []*v1.TaskStatusItem
	PageNo   int
	PageSize int
	Total    int
}
