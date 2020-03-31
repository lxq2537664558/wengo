//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"flag"
	"fmt"
	"github.com/showgo/conf"
	"github.com/showgo/csvdata"
	"github.com/showgo/global"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"time"
)

// 这里app 的初始化工作
func init() {

}

// 获取命令行启动
// 1. app相关的配置文件初始化
// 2. 设置app参数
func GetAppStart() {
	fmt.Println("App GetAppStart")
	
	//拉起宕机
	defer global.GrecoverToStd()
	
	proxy.InitProxy()
	csvdata.SetCsvPath(	proxy.PathModelPtr.CsvPath)
	csvdata.LoadPublicCsvData() // 读取公共的csv
	
	ParseCmd()//获取命令行
	OnAppStart()//解析完命令再启动对应程序
	proxy.AppWG.Add(1)
	// go AppRun()  // App 主要工作线程
	proxy.AppWG.Wait() // 等待app退出
	//CloseApp()   // 关闭app退出所有程序
}

// 程序启动获取命令行参数
func ParseCmd() {
	flag.IntVar(&proxy.AppID, "AppID", 0, "请输入app id")
	flag.Parse()
	for {
		proxy.AppConf = csvdata.GetAppconfPtr(proxy.AppID)
		if proxy.AppConf == nil {
			fmt.Println( "serverID 未找到")
		} else {
			break
		}
		time.Sleep(time.Second * 5)
	}
	fmt.Println( "ParseCmd success")
}

// 根据配置启动对应服务器
func OnAppStart() {
	initLog()
	proxy.InitKind()
	proxy.InitAppData(NewAppFactory(proxy.GetAppKind()))
	// 初始化app相关路径
	appBehavior = proxy.AppFactory.CreateAppBehavor()
	// 执行对应
	SetAppOpen()
	//读取控制台命令
	proxy.AppWG.Add(1)
	go ReadConsle()
}

func initLog() {
	logInit := &xlog.LogInitModel{
		ServerName:proxy.AppConf.App_name,
		LogsPath:proxy.PathModelPtr.LogsPath,
		Volatile:conf.VolatileModel,
	}
	xlog.NewXlog(logInit)
}


