package project_user

import (
	"context"
	"fmt"
	"wktline-server/internal/model"

	"wktline-server/api/project_user/v1"
	"wktline-server/internal/model/model_project"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelProjectUserByIds(ctx context.Context, req *v1.DelProjectUserByIdsReq) (res *v1.DelProjectUserByIdsRes, err error) {
	if err := service.Project().DelProjectUser(ctx, model_project.DelProjectUserInput{
		ProjectId: req.ProjectId,
		UserIds:   req.UserIds,
	}); err != nil {
		return nil, err
	}
	go func() {
		err = service.Message().CreateProjectActionUsersMsg(ctx, model.CreateProjectActionUsersMsgInput{
			ProjectId: req.ProjectId,
			UserIds:   req.UserIds,
			Title:     "你已被移出项目【%s】",
			MsgType:   "PROJECT_DEL_USER",
		})
		fmt.Println(err)
	}()
	return
}
