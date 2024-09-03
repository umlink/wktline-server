package v1

import "github.com/gogf/gf/v2/frame/g"

type UploadFileReq struct {
	g.Meta `path:"/common/upload" method:"post" mime:"multipart/form-data" summary:"上传文件（任意文件类型）" tags:"Upload"`
}
type UploadFileRes struct {
	Id   string `json:"id" v:"required" dc:"文件id"`
	Name string `json:"name" v:"required" dc:"文件名"`
	Type string `json:"type" v:"required" dc:"文件类型"`
	Url  string `json:"url" v:"required" dc:"文件url"`
}
