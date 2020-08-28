package Router

import (
	v2_0 "gin/Controllers/v2.0"
	"github.com/gin-gonic/gin"
)



func V2Router(Router *gin.RouterGroup) {
	v2 := Router.Group("v2.0")
	{
		v2.POST("/kk55kk", v2_0.Kk55kk)
	}

}