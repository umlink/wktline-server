package model_task

import (
	"time"

	v1 "wktline-server/api/task_labor_hour/v1"
)

type AddTaskLaborHourInput struct {
	Id          string
	TaskId      string
	ProjectId   string
	Date        time.Time
	Hour        float64
	Description string
	UserId      string
}

type UpdateTaskLaborHourInput struct {
	Id          string
	TaskId      string
	Date        time.Time
	Hour        float64
	Description string
}

type DelTaskLaborHourInput struct {
	Id     string
	TaskId string
	Hour   float64
}

type GetTaskLaborHourListInput struct {
	TaskId   string
	PageNo   int
	PageSize int
}
type GetTaskLaborHourListOutput struct {
	List     []*v1.TaskLaborHourItem
	Total    int
	PageSize int
	PageNo   int
}
