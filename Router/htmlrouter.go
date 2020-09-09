package Router

import (
	"gin/Controllers"
	"github.com/gin-gonic/gin"
)

func HtmlRouter(Router *gin.RouterGroup) {
	v1 := Router.Group("html")
	{
		v1.GET("/vlc", Controllers.HtmlIndex)
	}

}
