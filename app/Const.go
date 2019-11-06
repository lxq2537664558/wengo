// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.常量的定义
// 2.
// 3.
package app

type AppKind int

const(
	NONE 		   AppKind= 0			   //无类型
	Client  	   AppKind= 1			   //客户端
	LogonServer    AppKind= 2			   //登陆服
	GameServer     AppKind= 3			   //游戏服
	ChatServer     AppKind= 4			   //聊天服
	CenterServer   AppKind= 5	           //中心服务
)

// 整数变为AppKind
func ItoAppKind(val int) AppKind {
	switch (val) {
	case 0:
		return NONE
	case 1:
		return Client
	case 2:
		return LogonServer
	case 3:
		return GameServer
	case 4:
		return ChatServer
	case 5:
		return CenterServer
	default:
		return NONE
	}
}

func (ak AppKind) String() string {
	switch (ak) {
	case NONE:
		return "NONE"
	case Client: 
		return "Client"
	case LogonServer: 
		return "LogonServer"
	case GameServer: 
		return "GameServer"
	case ChatServer: 
		return "ChatServer"
	case CenterServer:
		return "CenterServer"
	default:         
		return "NONE"
	}
}
