// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserConfig is the golang structure for table user_config.
type UserConfig struct {
	UserId  string `json:"userId"  orm:"user_id"  description:"用户 id"`
	CallUrl string `json:"callUrl" orm:"call_url" description:"消息回调"`
}
