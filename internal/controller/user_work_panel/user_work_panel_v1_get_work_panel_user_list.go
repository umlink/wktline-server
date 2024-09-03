package user_work_panel

import (
	"context"
	"wktline-server/internal/model"
	"wktline-server/internal/service"

	"wktline-server/api/user_work_panel/v1"
)

func (c *ControllerV1) GetWorkPanelUserList(ctx context.Context, req *v1.GetWorkPanelUserListReq) (res *v1.GetWorkPanelUserListRes, err error) {
	out, err := service.WorkPanel().GetWorkPanelUserList(ctx, model.GetWorkPanelUserListInput{
		Keywords: req.Keywords,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
	})
	return &v1.GetWorkPanelUserListRes{
		List:     out.List,
		Total:    out.Total,
		PageNo:   out.PageNo,
		PageSize: out.PageSize,
	}, err
}
