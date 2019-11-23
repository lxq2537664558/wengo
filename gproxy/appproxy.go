/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package gproxy

import (
	. "../gmodel"
	"sync"
	"sync/atomic"
)

// App 相关数据存放
type AppProxy struct {
	AppInfo      *ServerAppInfo  // 服务器信息
	NetWorkInfo  *AppNetWorkInfo // 服务器网络信息
	AppKindArg   AppKind         // app类型 通过外部传递参数确定
	AppKindName  string          // app名称
	AppId        int32           // appID标识
	AppPath      string          // 路径
	AppWG        sync.WaitGroup  // app进程结束标志
	IsEndAppFlag int32           // app是否结束标志
}

// 创建AppProxy
func NewAppProxy() *AppProxy {
	appPro := new(AppProxy)
	return appPro
}

// 获取服务器信息
func (ap *AppProxy) GetServerAppInfo() *ServerAppInfo {
	return ap.AppInfo
}

// 获取服务器信息
func (ap *AppProxy) IsEnd() bool {
	return atomic.LoadInt32(&ap.IsEndAppFlag) == 1
}

// 获取服务器信息
func (ap *AppProxy) StopApp() {
	atomic.StoreInt32(&ap.IsEndAppFlag, 0)
}

// 获取服务器信息
func (ap *AppProxy) ReSartApp() {
	atomic.StoreInt32(&ap.IsEndAppFlag, 1)
}
