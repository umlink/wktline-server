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
	IUpload interface {
		UploadToQiNiu(ctx context.Context, in model.UploadFileInput) (res *model.UploadFileOutput, err error)
	}
)

var (
	localUpload IUpload
)

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for extendApi IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
