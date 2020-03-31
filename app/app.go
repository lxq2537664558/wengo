// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 程序最外层 ,这里给main的入口,以及整个进程退出的控制
// 包含程序的启动,停止
package app

import (
	"fmt"
	"github.com/showgo/app/appclient"
	"github.com/showgo/app/apploginsv"
	"github.com/showgo/model"
	"github.com/showgo/proxy"
	"github.com/showgo/xengine"
	"github.com/showgo/xlog"
)

var (
	EndFlag     *model.AtomicInt32FlagModel
	appBehavior xengine.ServerBehavior
)

// 逻辑app 主要工作线程
func AppRun() {
	defer xlog.RecoverToLog()
	defer proxy.AppWG.Done()
	for EndFlag.IsOpen() {
		// 运行app 逻辑
		if !appBehavior.OnUpdate()  {
		    break
		}
		
	}
	onCloseApp()
}

// 设置启动标志位
func SetAppOpen()  {
	EndFlag = model.NewAtomicInt32Flag()
	EndFlag.Open()
}

// 关闭app 程序
func CloseApp() {
	EndFlag.Close()
	onCloseApp()
}

// 执行关闭
func onCloseApp() {
	appBehavior.OnRelease()                // 进程结束
	proxy.RealseProxy()
	xlog.CloseLog() // 退出日志
	proxy.AppWG.Done()
	fmt.Println("App Close")
}

// app 逻辑参数根据服务器启动的参数创建对应的服务器工厂
func NewAppFactory(svKind model.AppKind) xengine.AppFactory {
	switch svKind {
	case model.APP_NONE:
		return nil
	case model.APP_Client:
		return new(appclient.ClientFactory)
	case model.APP_LoginServer: // 工厂
		return new(apploginsv.LoginServerFactory)
	case model.APP_GameServer:
		return nil
	case model.APP_MsgServer:
		return nil
	case model.APP_WorldServer:
		return nil
	default:
		return nil
	}
}

