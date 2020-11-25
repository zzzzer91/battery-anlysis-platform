package api

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/service"
	"battery-analysis-platform/pkg/jd"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		session := sessions.Default(c)
		userName := session.Get(constant.CookieKey)
		if userName == nil {
			c.JSON(200, jd.Err(""))
			return
		}
		s := service.LoginByCookieService{
			UserName: userName.(string),
		}
		jsonResponse(c, &s)
	} else if c.Request.Method == "POST" {
		s := service.LoginService{}
		// ShouldBind() 会检测是否满足设置的 bind 标签要求
		if err := c.ShouldBindJSON(&s); err != nil {
			c.AbortWithError(500, err)
			return
		}
		res, err := s.Do()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if res.Code == jd.SUCCESS {
			// 设置Session
			session := sessions.Default(c)
			session.Clear()
			session.Set(constant.CookieKey, s.UserName)
			session.Save()
		}
		c.JSON(200, res)
	} else {
		c.AbortWithError(500, errors.New("错误的 Request Method"))
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get(constant.CookieKey)
	if userName == nil {
		c.JSON(200, jd.Err(""))
		return
	}
	session.Clear()
	_ = session.Save()
	s := service.LogoutService{
		UserName: userName.(string),
	}
	jsonResponse(c, &s)
}
