package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Cors() gin.HandlerFunc { // HandlerFunc将gin中间件使用的处理程序定义为返回值
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"} // AllowMethods是允许客户端用于跨域请求的方法列表。默认值是简单方法(GET, POST, PUT, PATCH, DELETE, HEAD和OPTIONS)
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"} // AllowHeaders是客户端允许跨域请求使用的非简单报头的列表
	if gin.Mode() == gin.ReleaseMode { // Mode返回当前gin模式 // gin有三种模式：三种mode分别对应了不同的场景。在我们开发调试过程中，使用debug模式就可以了。在上线的时候，一定要选择release模式。而test可以用在测试场景中。
		// 生产环境需要配置跨域域名，否则403
		config.AllowOrigins = []string{"http://www.example.com"}
	}else {
		// 测试环境下模糊匹配本地开头的请求
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) { // regexp正则匹配，^是正则表达式匹配字符串开始位置 $是正则表达式匹配字符串结束位置
				return true // 正则表达式\d匹配一个数字，正则表达式\.用来匹配点字符
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	config.AllowCredentials = true // AllowCredentials指示请求是否可以包括用户凭证，如cookie、HTTP身份验证或客户端SSL证书。
	return cors.New(config) // New返回带有用户定义的自定义配置的位置中间件
}
