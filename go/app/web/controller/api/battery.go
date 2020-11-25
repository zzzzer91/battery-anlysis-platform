package api

import (
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func GetBatteryList(c *gin.Context) {
	s := service.GetBatteryListService{}
	if err := c.ShouldBindQuery(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	jsonResponse(c, &s)
}
