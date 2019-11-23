// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.
// 2.
// 3.
package app

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 实现app的生命周期
type ServerApp struct {
	// 服务器信息
	AppInfo *ServerAppInfo
	// 服务器网络信息
	NetWorkInfo AppNetWorkInfo
	// 配置接口
	conf  Confer
	appWg sync.WaitGroup
}

// 创建一个服务器
func NewServerApp() Lifer {
	sa := new(ServerApp)
	sa.conf = NewserverConf()
	return sa
}

// 获取服务器信息
func (sa *ServerApp) GetServerAppInfo() *ServerAppInfo {
	return sa.AppInfo
}

// app初始化工作
func (sa *ServerApp) init() bool{
	fmt.Println("协程数量 = ", runtime.NumGoroutine())
	if sa.conf == nil {
		return false
	}
	sa.conf.InitConf()
	
	return  true
}

func (sa *ServerApp) Start() {
	
	// 初始化成功才执行后面的方法
	if sa.init() {
		sa.appWg.Add(1)
		go sa.run() // App 主要工作线程
		sa.appWg.Wait() // 等待app退出
	}
	sa.Close()      // 关闭app退出所有程序
		
}

// 逻辑app 主要工作线程
func (sa *ServerApp) run() {
	
	fmt.Println("协程数量 = ", runtime.NumGoroutine())
	defer sa.appWg.Done()
	ch := time.After(time.Second * 3)
	select {
	case con := <-ch:
		fmt.Println(con)
		goto outfor
	}
outfor:
}

// 关闭app 程序
func (sa *ServerApp) Close() {
	// 进程结束
	// defer Gwp.Done()
	
}
