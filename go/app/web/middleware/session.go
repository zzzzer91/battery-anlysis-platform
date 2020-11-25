package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 0, Path: "/"})
	return sessions.Sessions("gin-session", store)
}
