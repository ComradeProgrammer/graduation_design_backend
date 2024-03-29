package app

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/app/db"
	"graduation_design/internal/app/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Run() {
	db.DBInit()
	//start a gin http server
	server := gin.Default()
	// use gin-session
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//set routes

	//login
	server.GET("/api/ping", handler.Ping)
	server.GET("/login", handler.Login)
	server.GET("/logout", handler.Logout)
	server.GET(config.REDIRECTROUTE, handler.Oauth)

	//projects
	server.GET("/api/projects", handler.GetProjects)
	server.GET("/api/projects/track", handler.TrackProject)
	server.GET("/api/projects/untrack", handler.UntrackProject)
	//milestones
	server.POST("/api/projects/milestone/create", handler.CreateMilestone)
	server.GET("/api/projects/milestone/all", handler.GetAllProjectMilestones)
	server.GET("/api/projects/milestone", handler.GetProjectMilestone)
	server.POST("/api/projects/milestone/edit", handler.EditMilestone)
	server.GET("/api/projects/milestone/delete", handler.DeleteProjectMilestone)
	//issues
	server.GET("/api/projects/issue/all", handler.GetAllProjectIssue)
	server.POST("/api/projects/milestone/issue/create", handler.CreateIssue)
	server.GET("/api/projects/milestone/issue/all", handler.GetAllIssues)
	server.GET("/api/projects/milestone/issue", handler.GetIssue)
	server.POST("/api/projects/milestone/issue/edit", handler.EditIssue)
	server.GET("/api/projects/milestone/issue/changestate", handler.ChangeIssueState)
	//code quality
	server.GET("/api/projects/quality", handler.GetCodeQuality)
	server.GET("/api/projects/job/log", handler.GetJobLog)
	server.GET("/api/projects/regex/all", handler.GetAllRegex)
	server.POST("/api/projects/regex/create", handler.CreateRegex)
	server.GET("/api/projects/regex/delete", handler.DeleteRegex)
	//analysis
	server.GET("/api/projects/statistic", handler.GetProjectOverviewStatistic)
	server.GET("/api/projects/statistic/user", handler.GetProjectUserStatistic)
	server.GET("/api/projects/statistic/language", handler.GetProjectLanguageStatistic)
	server.GET("/api/projects/statistic/commit", handler.GetProjectCommitStatistic)
	server.GET("/api/projects/statistic/activity", handler.GetProjectActivityStatistic)
	//tests
	server.GET("/api/test", handler.Test)
	//load frontend files
	// server.StaticFile("/","frontend/index.html")
	// server.Static("/assets", "frontend/assets")
	//run server

	server.Run(config.APPPORT)
}
