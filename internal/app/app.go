package app

import (
	"graduation_design/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	//start a gin http server 
	server := gin.Default()
	//set routes
	server.GET("/ping", handler.Ping)
	//run
	server.Run(":3333")
}
