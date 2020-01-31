/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
各个App的接口
*/

package app

import (
	"github.com/showgo/loginserver"
	"github.com/showgo/model"
)

type AppBehavior interface {
	// 程序启动
	StartApp()
	//初始化
	InitApp() bool
	// 程序运行
	RunApp()
	// 关闭
	QuitApp()
}



//app 逻辑参数根据服务器启动的参数创建对应的服务器接口
func NewAppBehavior(svKind  model.AppKind) AppBehavior {
	switch  svKind {
	case model.APP_NONE:
		return nil
	case model.APP_Client:
		return nil
	case model.APP_LoginServer: // 登录服逻辑接口
		return loginserver.NewLogionServer()
	case model.APP_GameServer:
		return nil
	case model.APP_ChatServer:
		return nil
	case model.APP_WorldServer:
		return nil
	default:
		return nil
	}
}