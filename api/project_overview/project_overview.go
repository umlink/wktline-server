// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package project_overview

import (
	"context"

	"wktline-server/api/project_overview/v1"
)

type IProjectOverviewV1 interface {
	GetProjectStatistics(ctx context.Context, req *v1.GetProjectStatisticsReq) (res *v1.GetProjectStatisticsRes, err error)
	GetProjectUserTaskStatusStatistics(ctx context.Context, req *v1.GetProjectUserTaskStatusStatisticsReq) (res *v1.GetProjectUserTaskStatusStatisticsRes, err error)
	GetProjectUserTaskTypeStatistics(ctx context.Context, req *v1.GetProjectUserTaskTypeStatisticsReq) (res *v1.GetProjectUserTaskTypeStatisticsRes, err error)
	GetProjectTaskTypeStat(ctx context.Context, req *v1.GetProjectTaskTypeStatReq) (res *v1.GetProjectTaskTypeStatRes, err error)
	GetProjectTaskStatusStat(ctx context.Context, req *v1.GetProjectTaskStatusStatReq) (res *v1.GetProjectTaskStatusStatRes, err error)
	GetProjectUserLaborHour(ctx context.Context, req *v1.GetProjectUserLaborHourReq) (res *v1.GetProjectUserLaborHourRes, err error)
	GetProjectUserTaskCount(ctx context.Context, req *v1.GetProjectUserTaskCountReq) (res *v1.GetProjectUserTaskCountRes, err error)
}
