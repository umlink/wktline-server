package cmd

import (
	"context"
	"wktline-server/internal/controller/auth"
	"wktline-server/internal/controller/message"
	"wktline-server/internal/controller/project"
	"wktline-server/internal/controller/project_invite"
	"wktline-server/internal/controller/project_overview"
	"wktline-server/internal/controller/task_priority"
	"wktline-server/internal/controller/test"
	"wktline-server/internal/controller/test_case"
	"wktline-server/internal/controller/user_work_panel"

	"wktline-server/internal/controller/task_labor_hour"
	"wktline-server/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"

	"wktline-server/internal/controller/login"
	"wktline-server/internal/controller/project_group"
	"wktline-server/internal/controller/project_user"
	"wktline-server/internal/controller/task"
	"wktline-server/internal/controller/task_actor"
	"wktline-server/internal/controller/task_attachment"
	"wktline-server/internal/controller/task_group"
	"wktline-server/internal/controller/task_operation_log"
	"wktline-server/internal/controller/task_status"
	"wktline-server/internal/controller/task_type"
	"wktline-server/internal/controller/upload"
	"wktline-server/internal/controller/user"
	"wktline-server/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi 自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}
			// *****************************************************************************************
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareCORS,
					// ghttp.MiddlewareHandlerResponse,
					service.Middleware().ResponseHandler,
					service.Middleware().Log,
				)
				group.Bind(
					test.NewV1(),
					login.NewV1(),
					auth.NewV1(),
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().Auth,
						service.Middleware().Ctx,
					)

					userCtr := user.NewV1()
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.GET("/user/getUseInfo", userCtr.GetUserInfo)
						group.GET("/user/getUserInfoById", userCtr.GetUserInfoById)
						group.POST("/user/getUserList", userCtr.GetUserList)
					})

					projectCtr := project.NewV1()
					projectGroupCtr := project_group.NewV1()
					projectUserCtr := project_user.NewV1()
					projectInviteCtr := project_invite.NewV1()
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.GET("/project/getProjectInfo", project.NewV1().GetProjectInfo)
						group.GET("/project/getProjectList", project.NewV1().GetProjectList)
						group.GET("/project/getProjectGroupList", projectGroupCtr.GetProjectGroupList)
						group.GET("/project/getProjectUser", projectUserCtr.GetProjectUser)
						group.GET("/project/getInviteInfo", projectInviteCtr.GetProjectInviteInfo)
						group.GET("/project/inJoinInvite", projectInviteCtr.InJoinInviteProject)
						group.GET("/project/getInviteList", projectInviteCtr.GetProjectMyInviteList)
					})

					taskGroupCtr := task_group.NewV1()
					taskTypeCtr := task_type.NewV1()
					taskStatusCtr := task_status.NewV1()
					taskPriorityCtr := task_priority.NewV1()
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.GET("/task/getTaskGroupList", taskGroupCtr.GetTaskGroupList)
						group.GET("/task/getTaskTypeList", taskTypeCtr.GetTaskTypeList)
						group.GET("/task/getTaskStatusList", taskStatusCtr.GetTaskStatusList)
						group.GET("/task/getTaskPriorityList", taskPriorityCtr.GetTaskPriorityList)
					})

					group.Bind(
						upload.NewV1(),
						task.NewV1(),
						task_actor.NewV1(),
						task_attachment.NewV1(),
						task_operation_log.NewV1(),
						task_labor_hour.NewV1(),
						user_work_panel.NewV1(),
						message.NewV1(),
						project_overview.NewV1(),
						test_case.NewV1(),
					)
					// 项目管理员才可修改（系统超管和管理员也有权限修改） *****************************************************************************************
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(
							service.Middleware().AuthProject,
						)
						group.Group("/", func(group *ghttp.RouterGroup) {
							group.POST("/project/updateProject", projectCtr.UpdateProject)
							group.POST("/project/addProjectUser", projectUserCtr.AddProjectUserByIds)
							group.POST("/project/delProjectUser", projectUserCtr.DelProjectUserByIds)
							group.POST("/project/updateProjectUserRole", projectUserCtr.UpdateProjectUserRole)
							group.GET("/project/genInviteCode", projectInviteCtr.GenProjectInviteCode)
							group.GET("/project/delInviteCode", projectInviteCtr.DelInviteCode)
						})
						group.Group("/", func(group *ghttp.RouterGroup) {
							group.POST("/task/createTaskGroup", taskGroupCtr.CreateTaskGroup)
							group.POST("/task/delTaskGroup", taskGroupCtr.DelTaskGroup)
							group.POST("/task/updateTaskGroup", taskGroupCtr.UpdateTaskGroup)
							group.POST("/task/createTaskType", taskTypeCtr.CreateTaskType)
							group.POST("/task/delTaskType", taskTypeCtr.DelTaskType)
							group.POST("/task/updateTaskType", taskTypeCtr.UpdateTaskType)
							group.POST("/task/createTaskStatus", taskStatusCtr.CreateTaskStatus)
							group.POST("/task/delTaskStatus", taskStatusCtr.DelTaskStatus)
							group.POST("/task/updateTaskStatus", taskStatusCtr.UpdateTaskStatus)
							group.POST("/task/updateTaskStatusSort", taskStatusCtr.UpdateTaskStatusSort)
							group.GET("/task/addTaskPriority", taskPriorityCtr.AddTaskPriority)
							group.GET("/task/updateTaskPriority", taskPriorityCtr.UpdateTaskPriority)
							group.GET("/task/delTaskPriority", taskPriorityCtr.DelTaskPriority)
						})
					})
					// 系统管理员可访问 *****************************************************************************************
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(
							service.Middleware().AuthAdmin,
						)
						group.Group("/", func(group *ghttp.RouterGroup) {
							group.POST("/user/createUser", userCtr.CreateUser)
							group.POST("/user/delUserById", userCtr.DelUserById)
							group.POST("/user/updateUserInfo", userCtr.UpdateUserInfo)
						})
						group.Group("", func(group *ghttp.RouterGroup) {
							group.POST("/project/createProject", projectCtr.CreateProject)
							group.POST("/project/delProject", projectCtr.DelProject)
							group.POST("/project/createProjectGroup", projectGroupCtr.CreateProjectGroup)
							group.POST("/project/delProjectGroup", projectGroupCtr.DelProjectGroup)
							group.POST("/project/updateProjectGroup", projectGroupCtr.UpdateProjectGroup)
						})
						// 系统超级管理员可访问 *****************************************************************************************
						group.Group("/", func(group *ghttp.RouterGroup) {
							group.Middleware(
								service.Middleware().AuthSuperAdmin,
							)
						})
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
