package model_task

import (
	v1 "wktline-server/api/task_priority/v1"
)

type AddTaskPriorityInput struct {
	Id    string
	Name  string
	Color string
}
type UpdateTaskPriorityInput struct {
	Id    string
	Name  string
	Color string
}

type GetTaskPriorityInput struct {
	PageNo   int
	PageSize int
}
type GetTaskPriorityOutput struct {
	List     []*v1.TaskPriorityItem
	PageNo   int
	PageSize int
	Total    int
}
