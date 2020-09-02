package main

import (
	"gin/Databases"
	"gin/Router"
)

func main() {
	defer Mysql.DB.Close()
	defer Mysql.RD.Close()
	Router.InitRouter()
}
