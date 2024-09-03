package user_work_panel

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/user_work_panel/v1"
)

func (c *ControllerV1) GetWorkLaborHourLogs(ctx context.Context, req *v1.GetWorkLaborHourLogsReq) (res *v1.GetWorkLaborHourLogsRes, err error) {
	// 针对第一次查询没有自己的情况，添加进工时面板
	myId := service.BizCtx().Get(ctx).User.Uid
	myCount, err := service.WorkPanel().GetWorkPanelForMyCount(ctx, myId)
	if err != nil {
		return nil, err
	}
	if myCount == 0 {
		if err = service.WorkPanel().AddUserForWorkPanel(ctx, model.AddWorkPanelUserInput{
			UserIds: []string{myId},
		}); err != nil {
			return nil, err
		}
	}
	out, err := service.WorkPanel().GetUsersWorkPanelList(ctx, model.GetUsersWorkPanelInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageSize:  req.PageSize,
		PageNo:    req.PageNo,
	})
	return &v1.GetWorkLaborHourLogsRes{
		List:     out.List,
		Total:    out.Total,
		PageNo:   out.PageNo,
		PageSize: out.PageSize,
	}, err
}
