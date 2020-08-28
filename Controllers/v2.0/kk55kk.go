package v2_0

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Kk55kk(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": "kk55kk",
	})
}
