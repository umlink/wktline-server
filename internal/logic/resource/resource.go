package resource

import (
	"context"

	"wktline-server/internal/dao"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

type sResource struct{}

func New() *sResource {
	return &sResource{}
}

func init() {
	service.RegisterResource(New())
}

func (s *sResource) CheckHasResourceByHash(ctx context.Context, hash string) (res *model.CheckHasResourceOutput, err error) {
	err = dao.Resource.Ctx(ctx).Where("hash", hash).Scan(&res)
	return res, err
}

func (s *sResource) CreateFileResource(ctx context.Context, in model.CreateResourceInput) (err error) {
	_, err = dao.Resource.Ctx(ctx).Data(in).Insert()
	return err
}

func (s *sResource) DelResourceById(ctx context.Context, id int64) (err error) {
	_, err = dao.Resource.Ctx(ctx).WherePri(id).Delete()
	return err
}

func (s *sResource) DelResourceByHash(ctx context.Context, hash string) (err error) {
	_, err = dao.Resource.Ctx(ctx).Where("hash", hash).Delete()
	return err
}
