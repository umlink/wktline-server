package user_work_panel

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/user_work_panel/v1"
)

func (c *ControllerV1) DelWorkPanelUser(ctx context.Context, req *v1.DelWorkPanelUserReq) (res *v1.DelWorkPanelUserRes, err error) {
	err = service.WorkPanel().DelUserForWorkPanel(ctx, model.DelWorkPanelUserInput{
		UserIds: req.UserIds,
	})
	return nil, err
}
