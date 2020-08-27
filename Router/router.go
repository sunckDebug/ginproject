package Router

import (
	"gin/Controllers"
	"gin/Middlewares"
	"gin/Sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化gin路由
	router := gin.Default()
	// 路由中间件中使用了自定的日志中间件
	router.Use(Middlewares.LoggerToFile())
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	// 路由分组
	v1 := router.Group("v1.0")
	// 要在路由组之前全局使用「jwt中间件」
	//v1.Use(Middlewares.JWTAuth())
	{
		v1.POST("/testinsert", Controllers.TestInsert)
		v1.POST("/add", Controllers.AddIndex)
		v1.POST("/sel", Controllers.SelIndex)
		v1.POST("/all", Controllers.AllIndex)
		v1.POST("/log", Controllers.LogIndex)
		v1.POST("/red", Controllers.RedIndex)
	}

	// 启动监听8080端口
	router.Run(":8080")
}