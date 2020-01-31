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
	APP_LogonServer         = 2 // 登陆服
	APP_GameServer          = 3 // 游戏服
	APP_ChatServer          = 4 // 聊天服
	APP_WorldServer         = 5 // 世界服
)
var Apps = [...]string{
	"none",
	"client",
	"logonserver",
	"gameserver",
	"chatserver",
	"worldserver",
}

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
	if ak >= APP_NONE  && ak <= APP_WorldServer {
		return  Apps[ak]
	}
	return  Apps[APP_NONE]
}
