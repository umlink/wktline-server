package upload

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"wktline-server/internal/model"
	"wktline-server/internal/service"
)

type sUpload struct{}

func New() *sUpload {
	return &sUpload{}
}

func init() {
	service.RegisterUpload(New())
}

func (s *sUpload) UploadToQiNiu(ctx context.Context, in model.UploadFileInput) (res *model.UploadFileOutput, err error) {
	// 裁剪文件后缀 - 文件类型
	suffixList := gstr.Split(in.FileName, ".")
	suffixStr := suffixList[len(suffixList)-1]
	if err != nil {
		return nil, err
	}
	// 七牛云上传配置
	accessKey := g.Cfg().MustGetWithEnv(ctx, `upload.accessKey`).String()
	secretKey := g.Cfg().MustGetWithEnv(ctx, `upload.secretKey`).String()
	bucket := g.Cfg().MustGetWithEnv(ctx, `upload.bucket`).String()
	server := g.Cfg().MustGetWithEnv(ctx, `upload.server`).String()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	// 开始上传
	key := in.Hash + "." + suffixStr
	if err = formUploader.PutFile(context.Background(), &ret, upToken, key, in.FilePath, &putExtra); err != nil {
		return nil, err
	}
	// 删除临时文件
	if err = gfile.Remove(in.FilePath); err != nil {
		return nil, err
	}

	return &model.UploadFileOutput{
		Name:       in.FileName,
		Type:       suffixStr,
		Url:        server + "/" + ret.Key,
		BucketHash: ret.Hash,
	}, nil
}
