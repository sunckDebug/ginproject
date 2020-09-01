package Controllers


import (
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
	if emailService.Username == ""{
		emailService.Username = Utils.USERNAME
	}
	if emailService.Content == ""{
		emailService.Content = "服务器发送的日志"
	}
	// 判断关键字段是否传递
	if emailService.Filename == ""{
		c.JSON(200, gin.H{"code": -1, "message": "filename字段不可为空"});return
	}


	// 发送邮件
	err2 := Utils.SendMail(Utils.USERNAME, emailService.Username, emailService.Content, emailService.Filename)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err2.Error()});return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "发送邮件成功", "data": ""});return
}



