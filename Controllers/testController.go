package Controllers

import (
	"fmt"
	"gin/Databases"
	"gin/Middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "strconv"
)


func SelIndex(c *gin.Context)  {

	// 获取请求的json参数定义了要传输的json字段如果不传字段则默认该类型的空值   int 0、string ""、
	type KDRespBody struct {
		Abc string `json:"abc"`
		Def int `json:"def"`
		Ghj struct {
			K1 int `json:"k_1"`
			K2 int `json:"k_2"`
			K3 struct {
				S1 int `json:"s_1"`
				S2 int `json:"s_2"`
			} `json:"k_3"`
		} `json:"ghj"`
		Qwe [5] int `json:"qwe"`
	}

	var reqInfo KDRespBody
	err := c.BindJSON(&reqInfo)
	if err != nil {
		//fmt.Println("------", err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	} else {
		//fmt.Println("++++++", reqInfo, reqInfo.Abc, reqInfo.Def)
	}

	if reqInfo.Abc == "" {
		c.JSON(200, gin.H{"code": 400, "message": "abc字段不可为空"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": reqInfo,
	})
	return
}

func AllIndex(c *gin.Context)  {
	//defer func() {
	//	err := recover() // recover内置函数，可以捕捉到异常
	//	if err != nil {  // 说明捕捉到错误
	//		// 请求时间
	//		datetime := time.Now().Format("2006-01-02 15:04:05")
	//		// 报错信息
	//		str1 := fmt.Sprintf("%s", err)
	//		// 请求路由
	//		reqUri := c.Request.RequestURI
	//		// 写Error级别的日志
	//		Middlewares.Logger().WithFields(logrus.Fields{
	//			"name": "zhh",
	//		}).Error(err)
	//		// 错误信息写入数据库记录
	//		error := Models.Error{DateTime: datetime, Body: str1, ReqUri: reqUri, Name: "zhh"}
	//		Mysql.DB.Create(&error)
	//	}
	//}()
	//
	//// 全查询
	//var email []Models.Email
	//Mysql.DB.Find(&email) // find product with code l1212
	//
	//// 带条件查询 模糊查询
	////Mysql.DB.Where("email LIKE ?", "%163%").Find(&email)
	//
	//for i := 0; i < len(email); i++ {
	//	fmt.Println(email[i].ID, email[i].UserID, email[i].Email, email[i].Subscribed)
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 1,
	//	"message": "success",
	//	"data": email,
	//})
	//return

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": "",
	})
}


func LogIndex(c *gin.Context)  {

	// 写Error级别的日志
	Middlewares.Logger().WithFields(logrus.Fields{
		"name": "hanyun",
	}).Error("记录一下日志", "Error")

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": "",
	})
	return
}

func RedIndex(c *gin.Context)  {

	// redis 字符串增删改查
	res0, err0 := Mysql.RD.Do("set", "key", "ZhangHaoHao").Result()
	fmt.Println("00000", res0, err0)

	res1, err1 := Mysql.RD.Do("get", "key").Result()
	fmt.Println("11111", res1, err1)

	res2, err2 := Mysql.RD.Do("del", "key").Result()
	fmt.Println("22222", res2, err2)

	res3, err3 := Mysql.RD.Do("get", "key").Result()
	fmt.Println("33333", res3, err3)


	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": "",
	})
	return
}