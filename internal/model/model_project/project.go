package model_project

import (
	"time"

	v1 "wktline-server/api/project/v1"
)

type CreateProjectInput struct {
	Id          string
	Name        string
	Description string
	HeaderImg   string
	OperatorId  string
	GroupId     string
	OwnerId     string
	CreatorId   string
	ShowType    string
}

type DelProjectInput struct {
	Id string
}

type UpdateProjectInput struct {
	Id          string
	Name        *string
	Description *string
	HeaderImg   *string
	GroupId     *string
	OwnerId     *string
	Status      *int
	ShowType    *string
	OperatorId  string
}

type GetProjectDetailInput struct {
	Id string
}
type GetProjectDetailOutput struct {
	Id           string
	Name         string
	Description  string
	HeaderImg    string
	GroupId      string
	GroupName    string
	OperatorId   string
	OperatorName string
	OwnerId      string
	OwnerName    string
	OwnerAvatar  string
	ShowType     string
	Status       int
	IsJoined     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GetProjectListInput struct {
	Keywords  *string
	GroupId   *string
	IsOwner   bool
	IsCreator bool
	IsJoined  bool
	ShowType  *string
	PageNo    int
	PageSize  int
}
type GetProjectListOutput struct {
	List     []*v1.ProjectInfoItem
	PageNo   int
	PageSize int
	Total    int
}
