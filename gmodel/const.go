// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.常量的定义
// 2.
// 3.
package gmodel

type AppKind int

const (
	APP_NONE        AppKind = 0 // 无类型
	APP_Client              = 1 // 客户端
	APP_LogonServer         = 2 // 登陆服
	APP_GameServer          = 3 // 游戏服
	APP_ChatServer          = 4 // 聊天服
	APP_WorldServer         = 5 // 世界服
)

// 整数变为AppKind
func ItoAppKind(val int) AppKind {
	switch (val) {
	case 0:
		return APP_NONE
	case 1:
		return APP_Client
	case 2:
		return APP_LogonServer
	case 3:
		return APP_GameServer
	case 4:
		return APP_ChatServer
	case 5:
		return APP_WorldServer
	default:
		return APP_NONE
	}
}

func (ak AppKind) ToString() string {
	switch (ak) {
	case APP_NONE:
		return "none"
	case APP_Client:
		return "client"
	case APP_LogonServer:
		return "logonserver"
	case APP_GameServer:
		return "gameserver"
	case APP_ChatServer:
		return "chatserver"
	case APP_WorldServer:
		return "worldserver"
	default:
		return "NONE"
	}
}
