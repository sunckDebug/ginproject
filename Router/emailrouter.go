package Router

import (
	"gin/Controllers"
	"github.com/gin-gonic/gin"
)



func EmailRouter(Router *gin.RouterGroup) {
	Email := Router.Group("email")
	{
		Email.POST("/create", Controllers.CreateIndex)
		Email.POST("/delete", Controllers.DeleteIndex)
		Email.POST("/update", Controllers.UpdateIndex)
		Email.POST("/retrieve", Controllers.RetrieveIndex)
		Email.POST("/send", Controllers.SendIndex)
		Email.GET("/file", Controllers.FileDownload)
		Email.POST("/multithreading", Controllers.Multithreading)
	}

}