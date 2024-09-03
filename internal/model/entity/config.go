// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	Id        string      `json:"id"        orm:"id"         description:""`
	Key       string      `json:"key"       orm:"key"        description:"配置文件 key"`
	Value     string      `json:"value"     orm:"value"      description:"JSON string"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
}
