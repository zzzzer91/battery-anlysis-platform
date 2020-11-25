package api

import (
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	s := service.CreateUserService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	jsonResponse(c, &s)
}

func GetUserList(c *gin.Context) {
	s := service.GetCommonUserListService{}
	jsonResponse(c, &s)
}

func UpdateUserInfo(c *gin.Context) {
	s := service.UpdateUserInfoService{
		UserName: c.Param("name"),
	}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	jsonResponse(c, &s)
}
