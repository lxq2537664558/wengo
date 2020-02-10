/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package proxy

import (
	"github.com/showgo/apploginsv"
	"github.com/showgo/model"
	"github.com/showgo/xengine"
	"sync"
)

// App 相关数据存放
type AppProxy struct {
	AppWG       sync.WaitGroup // app进程结束标志
	AppFactory    xengine.AppFactory
	AppKindArg    model.AppKind // app类型 通过外部传递参数确定
	AppServerName string        // app名称
}

// 创建AppProxy
func NewAppProxy() *AppProxy {
	appPro := new(AppProxy)
	return appPro
}


func (ap  *AppProxy)InitProxy(){
	ap.AppFactory = NewAppFactory(ap.AppKindArg)
}

// app 逻辑参数根据服务器启动的参数创建对应的服务器工厂
func NewAppFactory(svKind model.AppKind) xengine.AppFactory {
	switch svKind {
	case model.APP_NONE:
		return nil
	case model.APP_Client:
		return nil
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


func (ap  *AppProxy)RealseProxy(){

}

