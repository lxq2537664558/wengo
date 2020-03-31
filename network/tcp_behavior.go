/*
创建时间: 2020/2/21
作者: zjy
功能介绍:

*/

package network

import "net"

//waitGroup 操作
type WaitGrouper interface  {
	WaitGroupAddOne()
	WaitGroupDone()
	WaitGroupWait()
}

//

//tcpBehavior
type TcpBehavior interface {
	WaitGrouper
	RemoveConn(conn net.Conn)
	ReceiveData(tcpConn *TCPConn)       // 接收数据
}