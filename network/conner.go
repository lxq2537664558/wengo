package network

import (
	"net"
)

//连接接口
type Conner interface {
	ReadMsg() (error)
	WriteMsg(args ...[]byte) error
	WriteOneMsg(maincmd, subcmd uint16, msg []byte) error
	GetOneMsgByteArr(maincmd, subcmd uint16, msg []byte) ([]byte, error)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}

//网络事件向外传递
type NetWorkObserver interface {
	OnNetWorkConnect(conn Conner) error
	OnNetWorkClose(conn Conner) error
	OnNetWorkRead(msgdata *MsgData) error
}

// 消息接口
type HandlerNetWorkMsg func(conn Conner,maincmd,subcmd uint16,msgdata []byte) error