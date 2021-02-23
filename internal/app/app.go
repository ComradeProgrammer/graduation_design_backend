package app

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	//start a gin http server
	server := gin.Default()
	//set routes
	server.GET("/ping", handler.Ping)
	server.GET("/login", handler.Login)
	server.GET(config.REDIRECTROUTE, handler.Oauth)
	//run
	server.Run(config.APPPORT)
}
