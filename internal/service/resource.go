// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this extendApi file.
// ================================================================================

package service

import (
	"context"
	"wktline-server/internal/model"
)

type (
	IResource interface {
		CheckHasResourceByHash(ctx context.Context, hash string) (res *model.CheckHasResourceOutput, err error)
		CreateFileResource(ctx context.Context, in model.CreateResourceInput) (err error)
		DelResourceById(ctx context.Context, id int64) (err error)
		DelResourceByHash(ctx context.Context, hash string) (err error)
	}
)

var (
	localResource IResource
)

func Resource() IResource {
	if localResource == nil {
		panic("implement not found for extendApi IResource, forgot register?")
	}
	return localResource
}

func RegisterResource(i IResource) {
	localResource = i
}
