// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model/model_project"
)

type (
	IProject interface {
		CreateProjectGroup(ctx context.Context, in model_project.CreateGroupInput) (err error)
		DeleteGroupById(ctx context.Context, in model_project.DelGroupInput) (err error)
		UpdateGroup(ctx context.Context, in model_project.UpdateGroupInput) (err error)
		GetGroupDetail(ctx context.Context, in model_project.GetGroupDetailInput) (res *model_project.GetGroupDetailOutput, err error)
		GetGroupList(ctx context.Context, in model_project.GetGroupListInput) (res *model_project.GetGroupListOutput, err error)
		GenProjectInviteCode(ctx context.Context, in model_project.GenProjectInviteCodeInput) (code string, err error)
		GetProjectInviteInfo(ctx context.Context, in model_project.GetProjectInviteInfoInput) (res *model_project.GetProjectInviteInfoOutput, err error)
		InJoinProjectInvite(ctx context.Context, in model_project.InJoinProjectInviteInput) (err error)
		GetProjectMyInviteList(ctx context.Context, projectId string) (res *model_project.GetProjectMyInviteListOutput, err error)
		CheckProjectMyInvite(ctx context.Context, code string) (count int, err error)
		DelProjectMyInvite(ctx context.Context, code string) (err error)
		GetProjectUserTaskStatusStatistics(ctx context.Context, in model_project.GetProjectUserTaskStatusStatisticsInput) (res *model_project.GetProjectUserTaskStatusStatisticsOutput, err error)
		GetProjectUserTaskTypeStatistics(ctx context.Context, in model_project.GetProjectUserTaskTypeStatisticsInput) (res *model_project.GetProjectUserTaskTypeStatisticsOutput, err error)
		GetProjectTaskTypeStatistics(ctx context.Context, in model_project.GetProjectTaskTypeStatInput) (res *model_project.GetProjectTaskTypeStatOutput, err error)
		GetProjectTaskStatusStatistics(ctx context.Context, in model_project.GetProjectTaskStatusStatInput) (res *model_project.GetProjectTaskStatusStatOutput, err error)
		GetProjectStatistics(ctx context.Context, in model_project.GetProjectStatisticsInput) (res *model_project.GetProjectStatisticsOut, err error)
		GetProjectUserLaborHourStat(ctx context.Context, in model_project.GetProjectUserLaborHourStatInput) (res *model_project.GetProjectUserLaborHourStatOutput, err error)
		GetProjectUserTaskCountStat(ctx context.Context, in model_project.GetProjectUserTaskCountStatInput) (res *model_project.GetProjectUserTaskCountStatOutput, err error)
		CreateProject(ctx context.Context, in model_project.CreateProjectInput) (err error)
		DelProjectById(ctx context.Context, in model_project.DelProjectInput) (err error)
		UpdateProject(ctx context.Context, in model_project.UpdateProjectInput) (err error)
		GetProjectDetail(ctx context.Context, in model_project.GetProjectDetailInput) (project *model_project.GetProjectDetailOutput, err error)
		GetProjectCountFromGroupId(ctx context.Context, in string) (count int, err error)
		GetProjectList(ctx context.Context, in model_project.GetProjectListInput) (res *model_project.GetProjectListOutput, err error)
		GetProjectUserList(ctx context.Context, in model_project.GetProjectUserListInput) (res *model_project.GetProjectUserListOutput, err error)
		GetProjectUserDetail(ctx context.Context, in model_project.GetProjectUserDetailInput) (res *model_project.GetProjectUserDetailOutput, err error)
		GetProjectUserIds(ctx context.Context, projectId string) (ids []string, err error)
		AddProjectUserByIds(ctx context.Context, in model_project.AddProjectUserByIdsInput) (err error)
		AddProjectUser(ctx context.Context, in model_project.AddProjectUserInput) (err error)
		UpdateProjectUserRole(ctx context.Context, in model_project.UpdateProjectUserRoleInput) (err error)
		DelProjectUser(ctx context.Context, in model_project.DelProjectUserInput) (err error)
	}
)

var (
	localProject IProject
)

func Project() IProject {
	if localProject == nil {
		panic("implement not found for interface IProject, forgot register?")
	}
	return localProject
}

func RegisterProject(i IProject) {
	localProject = i
}
