package Mysql

import (
	"gin/Config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

var DB *gorm.DB

func init()  {
	mysqlIp := Config.ReadIni("mysql", "mysql_ip")
	mysqlPort := Config.ReadIni("mysql", "mysql_port")
	mysqlUser := Config.ReadIni("mysql", "mysql_user")
	mysqlPwd := Config.ReadIni("mysql", "mysql_pwd")
	mysqlDb := Config.ReadIni("mysql", "mysql_db")

	var err error
	//DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", mysqlUser + ":" +mysqlPwd + "@tcp(" + mysqlIp + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8&parseTime=True&loc=Local")

	// 以实现结构体名为非复数形式 默认是false
	DB.SingularTable(true)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}