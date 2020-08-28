package Controllers

import (
	Mysql "gin/Databases"
	"gin/Models"
	"gin/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加一条数据
func CreateIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}

	id, err := emailService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id})
	return
}


// 删除一条数据
func DeleteIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}

	id, err := emailService.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id})
	return
}


// 更新一条数据
func UpdateIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}

	id, err := emailService.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id})
	return
}


// 查询数据
func RetrieveIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}

	// 全查询
	var email []Models.Email
	Mysql.DB.Find(&email) // find product with code l1212

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": email})
	return
}
