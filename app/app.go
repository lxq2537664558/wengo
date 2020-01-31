// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 程序最外层 ,这里给main的入口,以及整个进程退出的控制
// 包含程序的启动,停止
package app

import (
	. "../proxy"
	"../xlog"
)

var (
	appBehavior AppBehavior
)

func G() {
	if err := recover(); err != nil {
		xlog.ErrorLogInterfaceParam(err)
	}
}


//1. app相关的配置文件初始化
//2. 设置app参数
func StartApp() {
	
	// 设置启动标志位
	AppPxy.EndFlag.Open()
	appBehavior = NewAppBehavior()
	appBehavior.StartApp()
	// 启动App
	AppPxy.AppWG.Add(1)
	go AppRun()         // App 主要工作线程
	AppPxy.AppWG.Wait() // 等待app退出
	CloseApp()          // 关闭app退出所有程序
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
