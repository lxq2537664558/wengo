// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.
// 2.
// 3.
package app

import (
	"sync"
	"runtime"
	"global"
)



//实现app的生命周期
type ServerApp struct {
	//服务器信息
	AppInfo *ServerAppInfo
	//服务器网络信息
	NetWorkInfo AppNetWorkInfo
	//配置接口
	conf Confer
	appWg sync.WaitGroup
}

//创建一个服务器
func NewServerApp() *ServerApp {
	sa := new(ServerApp)
	return sa
}

//获取服务器信息
func (sa *ServerApp)GetServerAppInfo() *ServerAppInfo {
	return  sa.AppInfo
}

//app初始化工作
func (sa *ServerApp)init()   {
	//sa.conf.LoadConf()
}

func (sa *ServerApp)Start()  {
	sa.init()
	sa.appWg.Add(1)
	go sa.run()    // App 主要工作线程
	sa.appWg.Wait()//等待app退出
	sa.Colse()     //关闭app
}


// 逻辑app 主要工作线程
func (sa *ServerApp)run()  {
	println( "协程数量 = " ,runtime.NumGoroutine())
	defer sa.appWg.Done()
	
	for{
		print()
	}
}

func (sa *ServerApp)Colse()  {
	//全局结束
	defer  Gwp.Done()

}