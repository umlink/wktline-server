package model_task

import (
	v1 "wktline-server/api/task_attachment/v1"
)

type AddTaskAttachmentInput struct {
	Id         string
	ProjectId  string
	TaskId     string
	ResourceId string
	CreatorId  string
}

type DelTaskAttachmentInput struct {
	Id string
}

type GetTaskAttachmentListInput struct {
	TaskId    string
	ProjectId *string
	CreatorId *string
	PageNo    int
	PageSize  int
}

type GetTaskAttachmentListOutput struct {
	List     []*v1.TaskAttachmentItem
	Total    int
	PageNo   int
	PageSize int
}
