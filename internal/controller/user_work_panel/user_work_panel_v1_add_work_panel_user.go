package user_work_panel

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/user_work_panel/v1"
)

func (c *ControllerV1) AddWorkPanelUser(ctx context.Context, req *v1.AddWorkPanelUserReq) (res *v1.AddWorkPanelUserRes, err error) {
	err = service.WorkPanel().AddUserForWorkPanel(ctx, model.AddWorkPanelUserInput{
		UserIds: req.UserIds,
	})
	return nil, err
}
