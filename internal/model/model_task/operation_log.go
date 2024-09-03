package model_task

import (
	v1 "wktline-server/api/task_operation_log/v1"
)

type CreateOperationLogInput struct {
	Id      string
	TaskId  string
	Type    v1.LogType
	Name    string
	Content string
	Desc    string
	UserId  string
}

type DelOperationLogInput struct {
	Id string
}

type GetOperationLogListInput struct {
	TaskId   string
	Type     v1.LogType
	PageNo   int
	PageSize int
}
type GetOperationLogListOutput struct {
	List     []*v1.TaskOperationLogItem
	Total    int
	PageNo   int
	PageSize int
}
