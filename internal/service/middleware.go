// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// AuthAdmin 管理员权限
		AuthAdmin(r *ghttp.Request)
		// AuthProject 可编辑项目授权
		AuthProject(r *ghttp.Request)
		// AuthSuperAdmin 超管权限
		AuthSuperAdmin(r *ghttp.Request)
		// AuthUser 普通登录用户
		AuthUser(r *ghttp.Request)
		// Context 自定义上下文对象
		Context(r *ghttp.Request)
		// Auth 前台系统权限控制，用户必须登录才能访问
		Auth(r *ghttp.Request)
		Ctx(r *ghttp.Request)
		// Log 日志打印
		Log(r *ghttp.Request)
		// ResponseHandler 返回处理中间件
		ResponseHandler(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
