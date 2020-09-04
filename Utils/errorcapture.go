package Utils

import (
	"bytes"
	"fmt"
	"gin/Middlewares"
	"github.com/sirupsen/logrus"
	"runtime"
)


// 错误捕获函数
func ErrorCapture() {
	err := recover() // recover内置函数，可以捕捉到异常
	if err != nil {  // 说明捕捉到错误
		e := printStackTrace(err)
		// 写Error级别的日志
		Middlewares.Logger().WithFields(logrus.Fields{
			"name": "zhh",
		}).Error(e)
	}
}


// 打印堆栈信息
func printStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
