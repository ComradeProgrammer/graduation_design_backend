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
	server.GET("/ping", handler.Ping)
	server.GET("/login", handler.Login)
	server.GET(config.REDIRECTROUTE, handler.Oauth)
	//projects
	server.GET("/projects", handler.GetProjects)
	server.GET("/projects/track", handler.TrackProject)
	server.GET("/projects/untrack", handler.UntrackProject)
	//milestones
	server.POST("/projects/milestone/create",handler.CreateMilestone)
	server.GET("/projects/milestone/all",handler.GetAllProjectMilestones)
	server.GET("/projects/milestone",handler.GetProjectMilestone)
	server.POST("/projects/milestone/edit",handler.EditMilestone)
	server.GET("/projects/milestone/delete",handler.DeleteProjectMilestone)
	//issues
	server.POST("/projects/milestone/issue/create",handler.CreateIssue)
	server.GET("/projects/milestone/issue/all",handler.GetAllIssues)
	server.GET("/projects/milestone/issue",handler.GetIssue)
	server.POST("/projects/milestone/issue/edit",handler.EditIssue)
	server.GET("/projects/milestone/issue/changestate",handler.ChangeIssueState)
	//tests
	server.GET("/test",handler.Test)
	//run server
	
	server.Run(config.APPPORT)
}
