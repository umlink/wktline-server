package model

import (
	v1 "wktline-server/api/user_work_panel/v1"
)

type AddWorkPanelUserInput struct {
	UserIds []string
}

type DelWorkPanelUserInput struct {
	UserIds []string
}

type GetUsersWorkPanelInput struct {
	StartTime string
	EndTime   string
	PageNo    int
	PageSize  int
}
type GetUsersWorkPanelOutput struct {
	List     []*v1.WorkLaborHourLogItem
	Total    int
	PageNo   int
	PageSize int
}

type GetWorkPanelUserListInput struct {
	Keywords string
	PageNo   int
	PageSize int
}
type GetWorkPanelUserListOutput struct {
	List     []*v1.WorkPanelUserItem
	PageNo   int
	PageSize int
	Total    int
}

type GetLaborHourByUserIdInput struct {
	UserId    string
	StartTime string
	EndTime   string
}
type GetLaborHourByUserIdOutput []*v1.WorkLaborHourDetailItem
