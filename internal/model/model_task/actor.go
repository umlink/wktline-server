package model_task

import (
	v1 "wktline-server/api/task_actor/v1"
)

type AddTaskActorInput struct {
	TaskId  string
	UserIds []string
	IsNew   bool
}

type DelTaskActorInput struct {
	TaskId  string
	UserIds []string
}

type GetTaskActorListInput struct {
	TaskId   string
	Keywords string
	PageNo   int
	PageSize int
}
type GetTaskActorListOutput struct {
	List     []*v1.TaskActorItem
	Total    int
	PageNo   int
	PageSize int
}

type FindTaskActorInput struct {
	TaskId string
	UserId string
}
