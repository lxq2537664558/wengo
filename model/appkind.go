// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.常量的定义
// 2.
// 3.
package model

type AppKind int

const (
	APP_NONE        AppKind = 0 // 无类型
	APP_Client              = 1 // 客户端
	APP_LoginServer         = 2 // 登陆服
	APP_GameServer          = 3 // 游戏服 各种场景处理
	APP_MsgServer           = 4 // 聊天服
	APP_WorldServer         = 5 // 世界服
)


var kindArr  =[...]AppKind{
	APP_NONE,
	APP_Client,
	APP_LoginServer,
	APP_GameServer,
	APP_MsgServer,
	APP_WorldServer,
}

// 整数变为AppKind
func ItoAppKind(val int) AppKind {
	if val >= 0 && val < len(kindArr) {
		return kindArr[val]
	}
	return APP_NONE
}

var appNames = [...]string{
	"none",
	"client",
	"loginsv",
	"gamesv",
	"msgsv",
	"worldsv",
}
func (ak AppKind) ToString() string {
	if ak >= APP_NONE  && ak < APP_WorldServer {
		return  appNames[ak]
	}
	return  appNames[APP_NONE]
}
