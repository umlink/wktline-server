// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model/model_task"
)

type (
	ITask interface {
		GetTaskActorList(ctx context.Context, in model_task.GetTaskActorListInput) (res *model_task.GetTaskActorListOutput, err error)
		AddTaskActor(ctx context.Context, in model_task.AddTaskActorInput) (err error)
		DelTaskActor(ctx context.Context, in model_task.DelTaskActorInput) (err error)
		FindTaskActor(ctx context.Context, in model_task.FindTaskActorInput) (count int, err error)
		AddTaskAttachment(ctx context.Context, in model_task.AddTaskAttachmentInput) (err error)
		DelTaskAttachment(ctx context.Context, in model_task.DelTaskAttachmentInput) (err error)
		GetTaskAttachmentList(ctx context.Context, in model_task.GetTaskAttachmentListInput) (res *model_task.GetTaskAttachmentListOutput, err error)
		CreateTaskGroup(ctx context.Context, in model_task.CreateTaskGroupInput) (err error)
		DelTaskGroup(ctx context.Context, in model_task.DelTaskGroupInput) (err error)
		UpdateTaskGroup(ctx context.Context, in model_task.UpdateTaskGroupInput) (err error)
		GetTaskGroup(ctx context.Context, in model_task.GetTaskGroupInput) (res *model_task.GetTaskGroupOutput, err error)
		AddTaskLaborHour(ctx context.Context, in model_task.AddTaskLaborHourInput) (err error)
		DelLaborHourById(ctx context.Context, in model_task.DelTaskLaborHourInput) (err error)
		UpdateTaskLaborHour(ctx context.Context, in model_task.UpdateTaskLaborHourInput) (err error)
		GetTaskLaborHourList(ctx context.Context, in model_task.GetTaskLaborHourListInput) (res *model_task.GetTaskLaborHourListOutput, err error)
		GenerateTaskUpdateLog(userId string, in model_task.UpdateTaskInput, task model_task.GetTaskDetailOutput) (err error)
		CreateTaskOperationLog(in model_task.CreateOperationLogInput) (err error)
		DelTaskOperationLog(ctx context.Context, in model_task.DelOperationLogInput) (err error)
		GetTaskOperationLogList(ctx context.Context, in model_task.GetOperationLogListInput) (res *model_task.GetOperationLogListOutput, err error)
		AddTaskPriority(ctx context.Context, in model_task.AddTaskPriorityInput) (err error)
		UpdateTaskPriority(ctx context.Context, in model_task.UpdateTaskPriorityInput) (err error)
		DelTaskPriority(ctx context.Context, id string) (err error)
		GetTaskPriority(ctx context.Context, in model_task.GetTaskPriorityInput) (res *model_task.GetTaskPriorityOutput, err error)
		// InitProjectTaskStatus 项目创建时同步调用一次 初始化项目下任务状态列表
		InitProjectTaskStatus(ctx context.Context, in []model_task.InitTaskStatusInput) (err error)
		CreateTaskStatus(ctx context.Context, in model_task.CreateTaskStatusInput) (err error)
		DelTaskStatus(ctx context.Context, in model_task.DelTaskStatusInput) (err error)
		GetTaskStatusById(ctx context.Context, statusId string) (res *model_task.GetTaskStatusByIdOutput, err error)
		UpdateTaskStatus(ctx context.Context, in model_task.UpdateTaskStatusInput) (err error)
		UpdateTaskStatusSort(ctx context.Context, in model_task.UpdateTaskStatusSortInput) (err error)
		GetTaskStatus(ctx context.Context, in model_task.GetTaskStatusInput) (res *model_task.GetTaskStatusOutput, err error)
		GetTaskCountByStatus(ctx context.Context, id string) (res int, err error)
		CheckUpdateForProjectUserAuth(ctx context.Context, projectId string) (role string, err error)
		DelTask(ctx context.Context, in model_task.DelTaskInput) (err error)
		GetTaskCountById(ctx context.Context, id string) (count int, err error)
		UpdateTask(ctx context.Context, in model_task.UpdateTaskInput) (err error)
		GetChildTaskList(in model_task.GetChildTaskListInput) (res *model_task.GetTaskListOutput, err error)
		CreateTask(ctx context.Context, in model_task.CreateTaskInput) (taskId string, err error)
		GetTaskDetail(in model_task.GetTaskDetailInput) (res *model_task.GetTaskDetailOutput, err error)
		GetTaskListByInterval(ctx context.Context, in model_task.GetTaskListByIntervalInput) (res *model_task.GetTaskListByIntervalOutput, err error)
		GetTaskList(in model_task.GetTaskListInput) (res *model_task.GetTaskListOutput, err error)
		InitProjectTaskType(ctx context.Context, in []model_task.InitTaskTypeInput) (err error)
		CreateTaskType(ctx context.Context, in model_task.CreateTaskTypeInput) (err error)
		DelTaskType(ctx context.Context, in model_task.DelTaskTypeInput) (err error)
		UpdateTaskType(ctx context.Context, in model_task.UpdateTaskTypeInput) (err error)
		GetTaskType(ctx context.Context, in model_task.GetTaskTypeInput) (res *model_task.GetTaskTypeOutput, err error)
	}
)

var (
	localTask ITask
)

func Task() ITask {
	if localTask == nil {
		panic("implement not found for interface ITask, forgot register?")
	}
	return localTask
}

func RegisterTask(i ITask) {
	localTask = i
}
