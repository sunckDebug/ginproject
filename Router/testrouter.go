package Router

import (
	"gin/Controllers"
	"gin/Middlewares"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	v1 := Router.Group("test")
	{
		v1.POST("/sel", Middlewares.JWTAuth(), Controllers.SelIndex)
		v1.POST("/all", Controllers.AllIndex)
		v1.POST("/log", Controllers.LogIndex)
		v1.POST("/red", Controllers.RedIndex)
	}

}
