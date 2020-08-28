package Router

import (
	v1_0 "gin/Controllers/v1.0"
	"gin/Middlewares"
	"github.com/gin-gonic/gin"
)

func V1Router(Router *gin.RouterGroup) {
	v1 := Router.Group("v1.0")
	{
		v1.POST("/testinsert", v1_0.TestInsert)
		v1.POST("/add", v1_0.AddIndex)
		v1.POST("/sel", Middlewares.JWTAuth(), v1_0.SelIndex)
		v1.POST("/all", v1_0.AllIndex)
		v1.POST("/log", v1_0.LogIndex)
		v1.POST("/red", v1_0.RedIndex)
	}

}
