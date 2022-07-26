package middleware

import (
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	// 用cookie存储session
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{ // 存储会话或会话存储的配置。字段是http的一个子集
		MaxAge: 1800,	// MaxAge=0表示没有指定“Max-Age”属性。MaxAge<0意味着现在删除cookie，相当于'Max-Age: 0'。MaxAge>0表示以秒为单位的Max-Age属性。
		Path:	"/",
		Secure:	false,
		HttpOnly: true,
	})
	return sessions.Sessions("camp-session", store) // cookie key是camp-session
}
