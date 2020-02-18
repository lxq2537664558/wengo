/*
创建时间: 2020/2/17
作者: zjy
功能介绍:

*/

package xlog

import (
	"fmt"
	"runtime"
)

var  LenStackBuf    = 4096
//拉起宕机标准输出
func GrecoverToStd() {
	if rec := recover(); rec != nil {
		if LenStackBuf > 0 {
			buf := make([]byte, LenStackBuf)
			l := runtime.Stack(buf, false)
			fmt.Printf("%v\n%s \n", rec, buf[:l])
		} else {
			fmt.Printf("%v\n", rec)
		}
	}
}

//拉起宕机日志输出
func GrecoverToLog() {
	if rec := recover(); rec != nil {
		if LenStackBuf > 0 {
			buf := make([]byte, LenStackBuf)
			l := runtime.Stack(buf, false)
			ErrorLog("app","%v\n%s", rec, buf[:l])
		} else {
			ErrorLog("app","%v", rec)
		}
	}
}
