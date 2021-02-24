package app

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/app/handler"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func Run() {
	//start a gin http server
	server := gin.Default()
	// use gin-session
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//set routes
	server.GET("/ping", handler.Ping)
	server.GET("/login", handler.Login)
	server.GET(config.REDIRECTROUTE, handler.Oauth)
	//run
	server.Run(config.APPPORT)
}
