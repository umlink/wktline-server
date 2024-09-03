// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user_work_panel

import (
	"context"
	
	"wktline-server/api/user_work_panel/v1"
)

type IUserWorkPanelV1 interface {
	AddWorkPanelUser(ctx context.Context, req *v1.AddWorkPanelUserReq) (res *v1.AddWorkPanelUserRes, err error)
	DelWorkPanelUser(ctx context.Context, req *v1.DelWorkPanelUserReq) (res *v1.DelWorkPanelUserRes, err error)
	GetWorkLaborHourLogs(ctx context.Context, req *v1.GetWorkLaborHourLogsReq) (res *v1.GetWorkLaborHourLogsRes, err error)
	GetWorkPanelUserList(ctx context.Context, req *v1.GetWorkPanelUserListReq) (res *v1.GetWorkPanelUserListRes, err error)
	GetWorkLaborHourByUserId(ctx context.Context, req *v1.GetWorkLaborHourByUserIdReq) (res *v1.GetWorkLaborHourByUserIdRes, err error)
}


