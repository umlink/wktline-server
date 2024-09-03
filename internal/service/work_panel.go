// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model"
)

type (
	IWorkPanel interface {
		GetWorkPanelForMyCount(ctx context.Context, uId string) (count int, err error)
		AddUserForWorkPanel(ctx context.Context, in model.AddWorkPanelUserInput) (err error)
		DelUserForWorkPanel(ctx context.Context, in model.DelWorkPanelUserInput) (err error)
		GetUsersWorkPanelList(ctx context.Context, in model.GetUsersWorkPanelInput) (res *model.GetUsersWorkPanelOutput, err error)
		GetWorkPanelByUserId(ctx context.Context, in model.GetLaborHourByUserIdInput) (res *model.GetLaborHourByUserIdOutput, err error)
		GetWorkPanelUserList(ctx context.Context, in model.GetWorkPanelUserListInput) (userList *model.GetWorkPanelUserListOutput, err error)
	}
)

var (
	localWorkPanel IWorkPanel
)

func WorkPanel() IWorkPanel {
	if localWorkPanel == nil {
		panic("implement not found for interface IWorkPanel, forgot register?")
	}
	return localWorkPanel
}

func RegisterWorkPanel(i IWorkPanel) {
	localWorkPanel = i
}
