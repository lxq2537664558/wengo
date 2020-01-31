//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"flag"
	"github.com/showgo/model"
	. "github.com/showgo/proxy"
	log "github.com/showgo/xlog"
	"os"
)



// 这里app 的初始化工作
func init() {
	InitConfig()
}

//初始化配置文件
func InitConfig()  {

}
// 获取命令行启动

//1. app相关的配置文件初始化
//2. 设置app参数
func GetStart() {
	getCommondLine()
	// 获取当前路径程序执行路径
	exepath, erro := os.Getwd()
	if erro != nil {
		log.DebugLog("app",erro.Error())
	}
	println(AppPxy.AppInfo.AppKindArg.ToString())
	PathPxy.SetAppPath(exepath)
	// 设置启动标志位
	AppPxy.EndFlag.Open()
	appBehavior = NewAppBehavior(AppPxy.AppInfo.AppKindArg)
	appBehavior.StartApp()
	
	
	// 启动App
	AppPxy.AppWG.Add(1)
	go AppRun()         // App 主要工作线程
	AppPxy.AppWG.Wait() // 等待app退出
	CloseApp()          // 关闭app退出所有程序
}

//程序启动获取命令行参数
func getCommondLine()  {
	println("App init")
	var intarg int
	flag.IntVar(&intarg, "appkind", 0, "请输入app类型")
	flag.Parse()
	if intarg == 0 {
		log.DebugLog("app","请输入app类型 -appkind > 0")
	}
	AppPxy.AppInfo.AppKindArg = model.ItoAppKind(intarg)
}