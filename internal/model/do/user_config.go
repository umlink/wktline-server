// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserConfig is the golang structure of table user_config for DAO operations like Where/Data.
type UserConfig struct {
	g.Meta  `orm:"table:user_config, do:true"`
	UserId  interface{} // 用户 id
	CallUrl interface{} // 消息回调
}
