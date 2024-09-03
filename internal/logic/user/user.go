package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
	"wktline-server/internal/dao"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) GetUserByUsernameAndPassword(ctx context.Context, in model.UserLoginInput) (user *model.UserLoginOut, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Username: in.Username,
		dao.User.Columns().Password: in.Password,
	}).Scan(&user)
	return user, err
}

func (s *sUser) CreateUser(ctx context.Context, in model.CreateUserInput) (res *model.CreateUserOutput, err error) {
	id, err := dao.User.Ctx(ctx).InsertAndGetId(in)
	return &model.CreateUserOutput{
		Id: gconv.String(id),
	}, err
}

func (s *sUser) CreateBatchUser(ctx context.Context, in model.CreateBatchUserInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).Batch(len(in)).Insert()
	return err
}

func (s *sUser) DeleteUserById(ctx context.Context, in model.DeleteUserInput) (err error) {
	_, err = dao.User.Ctx(ctx).WherePri(in.Id).Delete()
	return err
}

func (s *sUser) UpdateUserInfo(ctx context.Context, in model.UpdateUserInfoInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).OmitEmptyData().WherePri(in.Id).Update()
	return err
}

func (s *sUser) GetUserInfoById(ctx context.Context, in model.GetUserInfoInput) (user *model.GetUserInfoOutput, err error) {
	cacheKey := in.UserId
	cache, _ := gcache.Get(ctx, cacheKey)
	if !g.IsEmpty(cache) {
		var cRet model.GetUserInfoOutput
		if err = gconv.Struct(cache, &cRet); err == nil {
			return &cRet, nil
		}
	}
	if err = dao.User.Ctx(ctx).
		WherePri(in.UserId).
		FieldsEx("password,created_at,updated_at").
		Scan(&user); err != nil {
		return nil, err
	}
	// 这是缓存一小时
	if err = gcache.Set(ctx, in.UserId, user, time.Hour); err != nil {
		return nil, gerror.New("项目详情设置缓存失效")
	}
	return user, err
}

func (s *sUser) GetUserList(ctx context.Context, in model.GetUserListInput) (userList *model.GetUserListOutput, err error) {
	out := model.GetUserListOutput{
		PageNo:   in.PageNo,
		PageSize: in.PageSize,
	}
	d := dao.User.Ctx(ctx)
	if !g.IsEmpty(in.Keywords) {
		likePattern := `%` + in.Keywords + `%`
		d = d.Where(d.Builder().WhereLike("username", likePattern).WhereOrLike("nickname", likePattern))
	}
	if !g.IsEmpty(in.Role) {
		d = d.Where("role", in.Role)
	}
	if !g.IsEmpty(in.Status) {
		d = d.Where("status", in.Status)
	}
	if err = d.Page(in.PageNo, in.PageSize).ScanAndCount(&out.List, &out.Total, false); err != nil {
		return nil, err
	}
	return &out, err
}
