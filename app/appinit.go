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
	"github.com/showgo/model"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"os"
	"time"
)

// 这里app 的初始化工作
func init() {

}

// 获取命令行启动
// 1. app相关的配置文件初始化
// 2. 设置app参数
func GetStart() {
	fmt.Println("App GetStart")
	// 获取当前路径程序执行路径
	exepath, _ := os.Getwd()
	proxy.PathPxy.SetAppPath(exepath)
	proxy.PathPxy.InitProxy() //
	csvdata.SetCsvPath(	proxy.PathPxy.CsvPath)
	csvdata.LoadPublicCsvData()// 读取公共的csv
	getCommondLine() //获取命令行
	initConfig() //初始化配置文件
	initLog()

	
	// 设置启动标志位
	EndFlag = new(model.AtomicInt32FlagModel)
	EndFlag.Open()
	appBehavior = proxy.AppFactory.CreateAppBehavor()
	appBehavior.StartApp()
	// 启动App
	proxy.AppWG.Add(1)
	go AppRun()  // App 主要工作线程
	proxy.AppWG.Wait() // 等待app退出
	CloseApp()   // 关闭app退出所有程序
}

// 程序启动获取命令行参数
func getCommondLine() {
	flag.IntVar(&proxy.SververID, "serverID", 0, "请输入app类型")
	flag.Parse()
	for {
		proxy.SvConf = csvdata.GetServerconfPtr(proxy.SververID)
		if proxy.SvConf == nil {
			fmt.Println( "serverID 未找到")
		} else {
			break
		}
		time.Sleep(time.Second * 5)
	}
	proxy.InitKind()
	proxy.InitAppData(NewAppFactory(proxy.AppKindArg)) //初始化app相关路径
}

// 初始化配置文件
func initConfig() {
	conf.ReadIni(proxy.PathPxy.ConfIniPath)
}

func initLog() {
	logInit := &xlog.LogInitModel{
		ServerName:proxy.SvConf.ServerName,
		LogsPath:proxy.PathPxy.LogsPath,
		Volatile:conf.VolatileModel,
	}
	xlog.NewXlog(logInit)
}


