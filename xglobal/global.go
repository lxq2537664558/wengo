/*
创建时间: 2020/2/3
作者: zjy
功能介绍:

*/

package xglobal

import (
	"github.com/showgo/xlog"
	"github.com/showgo/conf"
	"runtime"
)


func Grecover() {
	if rec := recover(); rec != nil {
		if conf.LenStackBuf > 0 {
			buf := make([]byte, conf.LenStackBuf)
			l := runtime.Stack(buf, false)
			xlog.ErrorLog("xglobal","%v: %s", rec, buf[:l])
		} else {
			xlog.ErrorLog("xglobal","%v", rec)
		}
	}
}

