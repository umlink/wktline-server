package task

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"wktline-server/internal/consts"
	"wktline-server/internal/model/model_project"

	"wktline-server/api/task/v1"
	"wktline-server/internal/model/model_task"
	"wktline-server/internal/service"
)

func (c *ControllerV1) DelTask(ctx context.Context, req *v1.DelTaskReq) (res *v1.DelTaskRes, err error) {
	task, err := service.Task().GetTaskDetail(model_task.GetTaskDetailInput{
		Id: req.Id,
	})
	if g.IsEmpty(task) {
		return nil, gerror.New("任务不存在")
	}
	user := service.BizCtx().Get(ctx).User
	// 非本人：系统管理严、超管和项目管理严都可删除（直接跳转到删除操作），只有为普通用户进行判断
	if user.Role == consts.User {
		pUser, err := service.Project().GetProjectUserDetail(ctx, model_project.GetProjectUserDetailInput{
			ProjectId: req.ProjectId,
			UserId:    user.Uid,
		})
		if err != nil {
			return nil, gerror.New("查询用户归属项目失败")
		}
		if g.IsEmpty(pUser) {
			return nil, gerror.New("非本项目成员无权操作")
		}
		// 项目内的普通成员无法删除他人创建的任务
		if pUser.Role == consts.User {
			canDel := task.CreatorId == user.Uid
			// 非创建用户则继续检测是否是执行人 创建人和执行人可删除任务
			if !g.IsEmpty(task.HandlerId) && !canDel {
				canDel = task.HandlerId == user.Uid
			}
			if !canDel {
				return nil, gerror.New("非执行人和非创建人无权操作")
			}
		}
	}
	err = service.Task().DelTask(ctx, model_task.DelTaskInput{
		Id: req.Id,
	})
	return nil, err
}
