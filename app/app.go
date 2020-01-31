// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 程序最外层 ,这里给main的入口,以及整个进程退出的控制
// 包含程序的启动,停止
package app

import (
	."github.com/showgo/proxy"
	"github.com/showgo/xlog"
)

var (
	appBehavior AppBehavior
)

func G() {
	if err := recover(); err != nil {
		xlog.ErrorLogInterfaceParam(err)
	}
}




// 逻辑app 主要工作线程
func AppRun() {
	defer AppPxy.AppWG.Done()
	for AppPxy.EndFlag.IsOpen() {
		appBehavior.RunApp() // 运行app 逻辑
	}
}

// 关闭app 程序
func CloseApp() {
	// 进程结束
	appBehavior.QuitApp()
}
