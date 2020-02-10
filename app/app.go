// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 程序最外层 ,这里给main的入口,以及整个进程退出的控制
// 包含程序的启动,停止
package app

import (
	"github.com/showgo/model"
	."github.com/showgo/proxy"
	"github.com/showgo/xengine"
	"github.com/showgo/xglobal"
	"github.com/showgo/xlog"
)

var (
	EndFlag     *model.AtomicInt32FlagModel
	appBehavior xengine.AppBehavior
)

// 逻辑app 主要工作线程
func AppRun() {
	defer xglobal.Grecover()
	defer AppPxy.AppWG.Done()
	for EndFlag.IsOpen() {
		appBehavior.RunApp() // 运行app 逻辑
	}
	onCloseApp()
}

// 关闭app 程序
func CloseApp() {
	EndFlag.Close()
}

// 执行关闭
func onCloseApp() {
	appBehavior.QuitApp()                // 进程结束
	xlog.CloseLog(xlog.CloseType_nomarl) // 退出日志
	
	RealseProxy()
}
