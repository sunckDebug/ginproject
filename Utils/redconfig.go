package Utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)


func ReadIni(group string, key string) string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("文件读取错误", err)
		os.Exit(1)
	}

	value := fmt.Sprintf("%s", cfg.Section(group).Key(key))
	return value
}