package model_project

import v1 "wktline-server/api/project_invite/v1"

type GenProjectInviteCodeInput struct {
	ProjectId      string
	Deadline       string
	MaxInviteCount int
}

type GetProjectInviteInfoInput struct {
	Id string
}
type ProjectInviteDataItem struct {
	Id             string
	ProjectId      string
	JoinedCount    int
	MaxInviteCount int
}
type GetProjectInviteInfoOutput v1.GetProjectInviteInfoRes

type InJoinProjectInviteInput struct {
	Id        string
	ProjectId string
}
type GetProjectMyInviteListOutput v1.GetProjectMyInviteListRes
