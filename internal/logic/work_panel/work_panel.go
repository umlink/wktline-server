package work_panel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"strings"
	"time"
	"wktline-server/internal/dao"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
	"wktline-server/utility"
)

type sWorkPanel struct{}

func init() {
	service.RegisterWorkPanel(New())
}

func New() *sWorkPanel {
	return &sWorkPanel{}
}

func (s *sWorkPanel) GetWorkPanelForMyCount(ctx context.Context, uId string) (count int, err error) {
	return dao.UserWorkPanel.Ctx(ctx).Where("user_id", uId).Where("workmate_id", uId).Fields("id").Count()
}

func (s *sWorkPanel) AddUserForWorkPanel(ctx context.Context, in model.AddWorkPanelUserInput) (err error) {
	list := g.Slice{}
	for _, v := range in.UserIds {
		fmt.Println(time.Now())
		list = append(list, g.Map{
			"id":          utility.GenUniIdByGuid(),
			"workmate_id": v,
			"user_id":     service.BizCtx().Get(ctx).User.Uid,
		})
	}
	if _, err := dao.UserWorkPanel.Ctx(ctx).Data(list).Batch(len(list)).Insert(); err != nil {
		return err
	}
	return err
}

func (s *sWorkPanel) DelUserForWorkPanel(ctx context.Context, in model.DelWorkPanelUserInput) (err error) {
	_, err = dao.UserWorkPanel.Ctx(ctx).Where("user_id", service.BizCtx().Get(ctx).User.Uid).WhereIn("workmate_id", in.UserIds).Delete()
	return err
}

func (s *sWorkPanel) GetUsersWorkPanelList(ctx context.Context, in model.GetUsersWorkPanelInput) (res *model.GetUsersWorkPanelOutput, err error) {
	out := &model.GetUsersWorkPanelOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	myId := service.BizCtx().Get(ctx).User.Uid
	wd := dao.UserWorkPanel.Ctx(ctx).
		Fields("workmate_id").
		Where("user_id", myId).
		OrderAsc("id")
	if out.Total, err = wd.Count(); err != nil {
		return nil, err
	}
	userIds, err := wd.Page(in.PageNo, in.PageSize).Array()
	if err != nil {
		return nil, err
	}
	// whereIn 排序
	orderField := "field(id,'" + strings.Join(gconv.SliceStr(userIds), "'"+","+"'") + "')"
	ud := dao.User.Ctx(ctx).WhereIn("id", userIds).Order(orderField)
	ud = ud.Fields("id,avatar,username,nickname")
	if err = ud.ScanList(&out.List, "User"); err != nil {
		return nil, err
	}
	err = dao.TaskLaborHour.Ctx(ctx).
		Fields("hour,date,user_id").
		Where(dao.TaskLaborHour.Columns().UserId, gutil.ListItemValuesUnique(out.List, "User", "Id")).
		ScanList(&out.List, "Records", "User", "user_id:Id")
	return out, err
}

func (s *sWorkPanel) GetWorkPanelByUserId(ctx context.Context, in model.GetLaborHourByUserIdInput) (res *model.GetLaborHourByUserIdOutput, err error) {
	out := model.GetLaborHourByUserIdOutput{}
	d := g.Model("task_labor_hour th,task t").Where("th.user_id", in.UserId)
	d = d.Where("t.id=th.task_id")
	d = d.Where("th.date>=?", in.StartTime)
	d = d.Where("th.date<=?", in.EndTime)
	err = d.Fields("th.id,th.hour,th.description,th.date,th.task_id,t.name as task_name").Scan(&out)
	return &out, err
}

func (s *sWorkPanel) GetWorkPanelUserList(ctx context.Context, in model.GetWorkPanelUserListInput) (userList *model.GetWorkPanelUserListOutput, err error) {
	out := model.GetWorkPanelUserListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	myId := service.BizCtx().Get(ctx).User.Uid
	d := dao.User.Ctx(ctx)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.Where(d.Builder().WhereLike("username", likePattern).WhereOrLike("nickname", likePattern))
	}
	d = d.WhereNotIn("id", dao.UserWorkPanel.Ctx(ctx).Fields("workmate_id").Where("user_id", myId))
	d = d.Fields("id,username,avatar,nickname")
	if err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	return &out, err
}
