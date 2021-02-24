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
	//run server
	server.Run(config.APPPORT)
}
