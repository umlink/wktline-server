package model_project

import (
	"time"

	v1 "wktline-server/api/project_group/v1"
)

type CreateGroupInput struct {
	Id          string
	Name        string
	Description string
	OperatorId  string
}
type CreateGroupOutput struct {
	Id string
}

type DelGroupInput struct {
	Id string
}

type UpdateGroupInput struct {
	Id          string
	Name        *string
	Description *string
	OperatorId  string
}

type GetGroupListInput struct {
	Keywords string
	PageNo   int
	PageSize int
}

type GetGroupListOutput struct {
	List     []*v1.ProjectGroupItem
	PageNo   int
	PageSize int
	Total    int
}

type GetGroupDetailInput struct {
	Id string
}
type GetGroupDetailOutput struct {
	Id          string
	Name        string
	Description string
	UpdatedTime time.Time
}
