package Router

import (
	"gin/Middlewares"
	"gin/Sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化gin路由
	router := gin.Default()
	// 指明html加载文件目录
	router.LoadHTMLGlob("templates/*")
	// 指明静态文件的位置
	router.Static("static", "static")
	// 路由中间件中使用了自定的日志中间件
	router.Use(Middlewares.LoggerToFile())
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	// 方便统一添加路由组前缀 多服务器上线使用
	Group := router.Group("api")
	// 要在路由组之前全局使用「jwt中间件」
	//Group.Use(Middlewares.JWTAuth())
	// 注册TestRouter路由
	TestRouter(Group)
	// 注册EmailRouter路由
	EmailRouter(Group)
	// 注册HtmlRouter路由
	HtmlRouter(Group)
	// 启动监听8080端口
	router.Run(":8080")
}