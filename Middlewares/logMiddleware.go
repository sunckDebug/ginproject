package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

func Log()  {
	// 创建记录日志的文件
	f, _ := os.Create("logs/gin.log")
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}



//func main() {
//	// 记录到文件。
//
//	router := gin.Default()
//	router.Use(LoggerToFile())
//	router.GET("/", func(context *gin.Context) {
//		//Info级别的日志
//		Logger().WithFields(logrus.Fields{
//			"name": "hanyun",
//		}).Info("记录一下日志", "Info")
//		//Error级别的日志
//		Logger().WithFields(logrus.Fields{
//			"name": "hanyun",
//		}).Error("记录一下日志", "Error")
//		//Warn级别的日志
//		Logger().WithFields(logrus.Fields{
//			"name": "hanyun",
//		}).Warn("记录一下日志", "Warn")
//		//Debug级别的日志
//		Logger().WithFields(logrus.Fields{
//			"name": "hanyun",
//		}).Debug("记录一下日志", "Debug")
//	})
//	router.Run()
//}
