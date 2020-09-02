package Utils

import (
	"gin/Middlewares"
	"github.com/sirupsen/logrus"
)


// 错误捕获函数
func ErrorCapture() {
	err := recover() // recover内置函数，可以捕捉到异常
	if err != nil {  // 说明捕捉到错误
		// 写Error级别的日志
		Middlewares.Logger().WithFields(logrus.Fields{
			"name": "zhh",
		}).Error(err)
	}
}
