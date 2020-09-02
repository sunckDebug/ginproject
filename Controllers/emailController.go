package Controllers

import (
	"fmt"
	Mysql "gin/Databases"
	"gin/Models"
	"gin/Services"
	"gin/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 添加一条数据
func CreateIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()});return
	}

	id, err := emailService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()});return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id});return
}


// 删除一条数据
func DeleteIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()});return
	}

	id, err := emailService.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()});return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id});return
}


// 更新一条数据
func UpdateIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()});return
	}

	id, err := emailService.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": err.Error()});return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": id});return
}


// 查询数据
func RetrieveIndex(c *gin.Context)  {
	var emailService Services.Email

	err := c.ShouldBindJSON(&emailService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()});return
	}

	// 全查询
	var email []Models.Email
	Mysql.DB.Find(&email) // find product with code l1212

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": email});return
}


// 发送邮件
func SendIndex(c *gin.Context)  {
	// 接收接口传入数据
	var emailService Services.SendLog
	// 判断是否为json
	err1 := c.ShouldBindJSON(&emailService)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err1.Error()});return
	}


	// 如果为空赋默认值
	if(len(emailService.Toname) < 1){
		emailService.Toname = []string{Utils.USERNAME}
	}
	// 如果为空赋默认值
	if emailService.Content == ""{
		emailService.Content = "服务器发送的日志"
	}
	// 判断关键字段是否传递
	if emailService.Filename == ""{
		c.JSON(200, gin.H{"code": -1, "message": "filename字段不可为空"});return
	}


	// 发送邮件
	err2 := Utils.SendMail(Utils.USERNAME, emailService.Toname, emailService.Content, emailService.Filename)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err2.Error()});return
	}


	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "发送邮件成功", "data": ""});return
}


// 文件下载
func FileDownload(c *gin.Context){

	name := c.Query("name")

	// 判断关键字段是否传递
	if name == ""{
		c.JSON(200, gin.H{"code": -1, "message": "filename字段不可为空"});return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(name)
}
