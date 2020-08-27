package Utils

import (
	"gin/Middlewares"
	"github.com/sirupsen/logrus"
)


func ErrorCapture(fun func()) func() func() {
	wrapper := func() func(){
		defer func() {
			err := recover() // recover内置函数，可以捕捉到异常
			if err != nil {  // 说明捕捉到错误
				// 写Error级别的日志
				Middlewares.Logger().WithFields(logrus.Fields{
					"name": "zhh",
				}).Error(err)
			}
		}()
		return fun
	}
	return wrapper
}
