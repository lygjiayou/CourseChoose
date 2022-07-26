package server

import (
	"CourseChoose/api"
	"CourseChoose/server/middleware"
	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r :=  gin.Default()
	r.Use(middleware.Session("secret"))
	r.Use(middleware.Cors())
	/* Use的作用：使用连接全局中间件到路由器。例如，通过Use()附加的中间件将被包含在
	每个请求的处理程序链中。甚至404,405，静态文件…例如，这是记录器或错误管理中间件的正确位置
	 */
	// 路由
	g := r.Group("/api/v1")
	{
		// 权限鉴定
		//admin := g.Group("")
		/* Group创建新的路由器组。您应该添加所有具有公共中间件或相同路径前缀的路由。
		* 例如，可以对所有使用公共中间件进行授权的路由进行分组。
		*/
		//admin.Use(middleware)
		//{
		//
		//}
		// 成员管理（无鉴权）
		g.POST("/member/create", api.CreateMember)
	}
	return r
}