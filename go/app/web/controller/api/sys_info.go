package api

import (
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func GetSysInfo(c *gin.Context) {
	s := service.GetSysInfoService{}
	jsonResponse(c, &s)
}
