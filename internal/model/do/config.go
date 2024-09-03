// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table config for DAO operations like Where/Data.
type Config struct {
	g.Meta    `orm:"table:config, do:true"`
	Id        interface{} //
	Key       interface{} // 配置文件 key
	Value     interface{} // JSON string
	UpdatedAt *gtime.Time //
	CreatedAt *gtime.Time //
}
