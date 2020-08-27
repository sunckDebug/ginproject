package main

import (
	"gin/Databases"
	"gin/Middlewares"
	"gin/Router"
)

func main() {
	Middlewares.Log()
	defer Mysql.DB.Close()
	defer Mysql.RD.Close()
	Router.InitRouter()
}
