package model_project

import (
	v1 "wktline-server/api/project_user/v1"
)

type AddProjectUserByIdsInput struct {
	ProjectId string
	UserIds   []string
}

type AddProjectUserInput struct {
	Id        string
	ProjectId string
	UserId    string
	Role      string
}

type UpdateProjectUserRoleInput struct {
	ProjectId string
	UserId    string
	Role      string
}

type DelProjectUserInput struct {
	ProjectId string
	UserIds   []string
}

type GetProjectUserListInput struct {
	ProjectId string
	Keywords  string
	Role      string
	PageNo    int
	PageSize  int
}
type GetProjectUserListOutput struct {
	List     []*v1.ProjectUserItem
	Total    int
	PageNo   int
	PageSize int
}

type GetProjectUserDetailInput struct {
	ProjectId string
	UserId    string
}

type GetProjectUserDetailOutput struct {
	Id     string
	UserId string
	Role   string
}
