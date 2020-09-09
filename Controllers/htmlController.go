package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加一条数据
func HtmlIndex(c *gin.Context)  {
	// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
	c.HTML(http.StatusOK, "vlc.html", nil)
}