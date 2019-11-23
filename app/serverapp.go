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
)

// 实现app的生命周期
type ServerApp struct {
	// 配置接口
	conf  Confer

}

// 创建一个服务器
func NewServerApp() Lifer {
	sa := new(ServerApp)
	sa.conf = NewserverConf()
	return sa
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
		// 设置启动标志位
		AppPxy.ReSartApp()
		AppPxy.AppWG.Add(1)
		go sa.run() // App 主要工作线程
		AppPxy.AppWG.Wait() // 等待app退出
	}
	sa.Close()      // 关闭app退出所有程序
		
}

// 逻辑app 主要工作线程
func (sa *ServerApp) run() {
	
	fmt.Println("协程数量 = ", runtime.NumGoroutine())
	defer AppPxy.AppWG.Done()
	for AppPxy.IsEnd() {
	
	}
}

// 关闭app 程序
func (sa *ServerApp) Close() {
	// 进程结束

	
}
