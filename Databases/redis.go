package Mysql

import (
	"fmt"
	"gin/Config"
	"github.com/go-redis/redis"
)

var RD *redis.Client

func init() {
	redisIp := Config.ReadIni("redis", "redis_ip")
	redisPort := Config.ReadIni("redis", "redis_port")

	RD = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(redisIp + ":" + redisPort),
		Password: "",
		DB:       0,
	})

	_, err := RD.Ping().Result()
	if err != nil {
		fmt.Println("redis ping error", err.Error())
	}
}