package api

import (
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func UpdateSelfPassword(c *gin.Context) {
	val, _ := c.Get("user")
	user := val.(*model.User)
	s := service.UpdateUserPasswordService{
		UserName: user.Name,
	}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	jsonResponse(c, &s)
}
