package web

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/controller"
	"battery-analysis-platform/app/web/middleware"
	"github.com/gin-gonic/gin"
)

func Run() error {
	gin.SetMode(conf.App.Gin.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(conf.App.Gin.SecretKey))
	controller.Register(r)
	return r.Run(conf.App.Gin.HttpAddr)
}
