package network

import (
	"fmt"
	"github.com/showgo/csvdata"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	sync.Mutex
	Addr            string
	ConnNum         int
	ConnectInterval time.Duration
	PendingWriteNum int
	AutoReconnect   bool
	netObserver NetWorkObserver
	NewAgent        func(*TCPConn) Agent
	conns           ConnSet
	wg              sync.WaitGroup
	closeFlag       bool
	appConf            *csvdata.Appconf
}

//创建tcp Sever服务器
func NewTCPClient(netobs NetWorkObserver,appconf *csvdata.Appconf ) *TCPClient {
	if appconf == nil {
		xlog.WarningLog(proxy.GetSecenName(),"server conf is nil")
		return nil
	}
	
	client := new(TCPClient)
	client.netObserver = netobs
	client.appConf = appconf
	return client
}

func (client *TCPClient) Start() {
	client.init()

	for i := 0; i < client.ConnNum; i++ {
		client.wg.Add(1)
		go client.connect()
	}
}

func (client *TCPClient) init() {
	client.Lock()
	defer client.Unlock()

	if client.ConnNum <= 0 {
		client.ConnNum = 1
		xlog.DebugLog(proxy.GetSecenName(),"invalid ConnNum, reset to %v", client.ConnNum)
	}
	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		xlog.DebugLog(proxy.GetSecenName(),"invalid ConnectInterval, reset to %v", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		xlog.DebugLog(proxy.GetSecenName(),"invalid PendingWriteNum, reset to %v", client.PendingWriteNum)
	}
	if client.NewAgent == nil {
		xlog.WarningLog(proxy.GetSecenName(),"NewAgent must not be nil")
	}
	if client.conns != nil {
		xlog.WarningLog(proxy.GetSecenName(),"client is running")
	}

	client.conns = make(ConnSet)
	client.closeFlag = false
	client.AutoReconnect = true

	// msg parser
	msgParser = NewMsgParser(client.appConf.Msglen_size,client.appConf.Max_msglen)
	client.Addr = fmt.Sprintf("%s:%s",client.appConf.Out_addr,client.appConf.Out_prot)
	xlog.DebugLog(proxy.GetSecenName(),"client.Addr connet %v ", client.Addr)
}

func (client *TCPClient) dial() net.Conn {
	for {
		conn, err := net.Dial("tcp", client.Addr)
		if err == nil || client.closeFlag {
			return conn
		}

		xlog.DebugLog(proxy.GetSecenName(),"connect to %v error: %v", client.Addr, err)
		time.Sleep(client.ConnectInterval)
		continue
	}
}

func (client *TCPClient) connect() {
	defer client.wg.Done()

	conn := client.dial()
	if conn == nil {
		client.reConnect() //连接不上要重新连接
		return
	}
	
	if !client.addConn(conn) {
		return
	}

	tcpConn := newTCPConn(conn, client.netObserver,client.appConf)
	client.wg.Add(1)
	go client.ReceiveData(tcpConn) //读取远端数据
}

//添加链接信息
func (client *TCPClient) addConn(conn net.Conn) bool  {
	client.Lock()
	if client.closeFlag {
		client.Unlock()
		conn.Close()
		return false
	}
	client.conns[conn] = struct{}{}
	client.Unlock()
	xlog.DebugLog(proxy.GetSecenName(),"连接",conn.RemoteAddr(),"成功")
	return true
}

// 连接中读取数据
func (client *TCPClient) ReceiveData(tcpConn *TCPConn) {
	defer client.wg.Done() //关闭连接要释放
	
	for {
		err := tcpConn.ReadMsg()
		if err != nil { // 这里读到错误消息,关闭
			xlog.WarningLog(proxy.GetSecenName(), "read message: ", err)
			break // 关闭连接
		}
	}
	// cleanup
	client.RemoveConn(tcpConn)
	
	client.reConnect()
}

func (client *TCPClient) reConnect() {
	// 进行重连
	if client.AutoReconnect {
		xlog.DebugLog(proxy.GetSecenName(),"重新连接")
		time.Sleep(client.ConnectInterval)
		client.wg.Add(1)
		go client.connect()
	}
}

func (client *TCPClient) RemoveConn(tcpConn *TCPConn) {
	tcpConn.Close()
	client.Lock()
	delete(client.conns, tcpConn.conn)
	client.Unlock()
	xlog.DebugLog(proxy.GetSecenName(),"关闭远程连接")
}

func (client *TCPClient) Close() {
	client.Lock()
	client.closeFlag = true
	for conn := range client.conns {
		conn.Close()
	}
	client.conns = nil
	client.Unlock()

	client.wg.Wait()
}
