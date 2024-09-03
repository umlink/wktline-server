package consts

import (
	"wktline-server/internal/model/model_task"
	"wktline-server/utility"
)

// 此处三个状态贯穿整个项目，不可修改！！！，只可添加，不可删减。
const (
	TaskPadding    = "PADDING"
	TaskProcessing = "PROCESSING"
	TaskDone       = "DONE"
)

func GetDefaultTaskStatusList(projectId string, operatorId string) []model_task.InitTaskStatusInput {
	return []model_task.InitTaskStatusInput{
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "待处理",
			Sort:       1,
			Color:      "#DAA520",
			Enum:       TaskPadding,
			Default:    1,
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "处理中",
			Sort:       2,
			Color:      "#722ed1",
			Enum:       TaskProcessing,
			Default:    0,
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "已完成",
			Sort:       3,
			Color:      "#13a8a8",
			Enum:       TaskDone,
			Default:    0,
		},
	}
}

func GetDefaultTaskTypeList(projectId string, operatorId string) []model_task.InitTaskTypeInput {
	return []model_task.InitTaskTypeInput{
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "需求",
			Color:      "#DAA520",
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "开发",
			Color:      "#722ed1",
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "UI设计",
			Color:      "#13a8a8",
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "测试",
			Color:      "#13a8a8",
		},
		{
			Id:         utility.GenUniIdByGuid(),
			ProjectId:  projectId,
			OperatorId: operatorId,
			Name:       "提效",
			Color:      "#13a8a8",
		},
	}
}
