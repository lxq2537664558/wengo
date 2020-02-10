//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"flag"
	"fmt"
	"github.com/showgo/model"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"os"
)

// 这里app 的初始化工作
func init() {

}

// 获取命令行启动
// 1. app相关的配置文件初始化
// 2. 设置app参数
func GetStart() {
	fmt.Println("App GetStart")
	getCommondLine()
	initConfig()
	initLog()
	// 获取当前路径程序执行路径
	exepath, erro := os.Getwd()
	if erro != nil {
		xlog.DebugLog("app", erro.Error())
	}
	proxy.InitProxy() // 初始化代理包
	proxy.PathPxy.SetAppPath(exepath)
	// 设置启动标志位
	EndFlag.Open()
	appBehavior = proxy.AppPxy.AppFactory.CreateAppBehavor()
	appBehavior.StartApp()
	// 启动App
	proxy.AppPxy.AppWG.Add(1)
	go AppRun()  // App 主要工作线程
	proxy.AppPxy.AppWG.Wait() // 等待app退出
	CloseApp()   // 关闭app退出所有程序
}

// 程序启动获取命令行参数
func getCommondLine() {
	var intarg int
	flag.IntVar(&intarg, "appkind", 0, "请输入app类型")
	flag.Parse()
	for {
		proxy.AppPxy.AppKindArg = model.ItoAppKind(intarg)
		if proxy.AppPxy.AppKindArg == model.APP_NONE {
			xlog.DebugLog("app", "请输入app类型 -appkind > 0")
		} else {
			break
		}
	}
	proxy.AppPxy.AppServerName = proxy.AppPxy.AppKindArg.ToString()
}

// 初始化配置文件
func initConfig() {

}

func initLog() {
	logInit := &xlog.LogInitModel{
		proxy.AppPxy.AppServerName,
		proxy.PathPxy.LogsPath,
		xlog.VolatileLogModel{
			500,
			true,
			7,
		},
	}
	xlog.NewXlog(logInit)
}


