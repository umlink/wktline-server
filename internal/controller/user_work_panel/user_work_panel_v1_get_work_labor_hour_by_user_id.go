package user_work_panel

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/user_work_panel/v1"
)

func (c *ControllerV1) GetWorkLaborHourByUserId(ctx context.Context, req *v1.GetWorkLaborHourByUserIdReq) (res *v1.GetWorkLaborHourByUserIdRes, err error) {
	out, err := service.WorkPanel().GetWorkPanelByUserId(ctx, model.GetLaborHourByUserIdInput{
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	return (*v1.GetWorkLaborHourByUserIdRes)(out), err
}
