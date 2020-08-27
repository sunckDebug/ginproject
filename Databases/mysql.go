package Mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

var DB *gorm.DB

func init()  {
	var err error
	DB, err = gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}