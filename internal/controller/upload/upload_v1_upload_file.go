package upload

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"wktline-server/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"wktline-server/api/upload/v1"
	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	uploadPath := g.Cfg().MustGetWithCmd(gctx.New(), `upload.path`).String()
	var tempFileName string
	if tempFileName, err = file.Save(uploadPath, true); err != nil {
		return
	}
	fileSize := file.Size
	filePath := uploadPath + tempFileName
	hash, err := GetFileMD5(filePath)
	if err != nil {
		return nil, err
	}
	out, err := service.Upload().UploadToQiNiu(ctx, model.UploadFileInput{
		FilePath: filePath,
		FileName: file.Filename,
		Size:     file.Size,
		Hash:     hash,
	})
	// 根据文件生成key - md5 唯一性
	if err != nil {
		return nil, err
	}
	resource, err := service.Resource().CheckHasResourceByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if !g.IsEmpty(resource) {
		return &v1.UploadFileRes{
			Id:   resource.Id,
			Url:  resource.Url,
			Type: resource.Type,
			Name: resource.Name,
		}, err
	}
	fileId := utility.GenUniIdByGuid()
	if err != nil {
		return nil, err
	}
	err = service.Resource().CreateFileResource(ctx, model.CreateResourceInput{
		Id:         fileId,
		Name:       out.Name,
		Url:        out.Url,
		Type:       out.Type,
		Size:       fileSize,
		Hash:       hash,
		BucketHash: out.BucketHash,
		CreatorId:  service.BizCtx().Get(ctx).User.Uid,
	})
	return &v1.UploadFileRes{
		Id:   fileId,
		Url:  out.Url,
		Name: out.Name,
		Type: out.Type,
	}, err
}
