package main

import (
	"graduation_design/internal/app"
	"graduation_design/internal/pkg/logs"
)

func main() {
	logs.Info("app start")
	app.Run()
}
