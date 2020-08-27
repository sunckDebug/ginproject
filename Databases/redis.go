package Mysql

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RD *redis.Client

func init() {
	RD = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("127.0.0.1:6379"),
		Password: "",
		DB:       0,
	})

	_, err := RD.Ping().Result()
	if err != nil {
		fmt.Println("redis ping error", err.Error())
	}
}