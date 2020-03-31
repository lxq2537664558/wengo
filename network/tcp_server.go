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

type TCPServer struct {
	ln          net.Listener //服务器监听对象
	conns       ConnSet      //连接的对象
	connetsize  int          //已经连接的数量
	mutexConns  sync.RWMutex
	wgLn        sync.WaitGroup
	wgConns     sync.WaitGroup
	netObserver NetWorkObserver
	appConf   *csvdata.Appconf
}




//创建tcp Sever服务器
func NewTcpServer(netobs NetWorkObserver,appconf *csvdata.Appconf) *TCPServer {
	if appconf == nil {
		xlog.WarningLog(proxy.GetSecenName(),"server conf is nil")
		return nil
	}
	tcpsv := new(TCPServer)
	tcpsv.netObserver = netobs
	tcpsv.appConf = appconf
	return tcpsv
}

func (server *TCPServer) Start() {
	xlog.DebugLog(proxy.GetSecenName(),"TCPServer start")
	server.init()
	go server.run()
}

func (server *TCPServer) init() {
	ln, err := net.Listen("tcp",fmt.Sprintf(":%s",server.appConf.Out_prot))
	xlog.DebugLog(proxy.GetSecenName(),"TCPServer listen Addr:%v", ln.Addr())
	if err != nil {
		xlog.DebugLog(proxy.GetSecenName(),"%v", err)
	}

	if server.appConf.Max_connect  <= 0 {
		server.appConf.Max_connect = 10000
		xlog.WarningLog(proxy.GetSecenName(),"invalid MaxConnNum, reset to %v", server.appConf.Max_connect)
	}
	if server.appConf.Write_cap_num <= 0 {
		server.appConf.Write_cap_num  = 200
		xlog.WarningLog(proxy.GetSecenName(),"invalid PendingWriteNum, reset to %v", server.appConf.Write_cap_num)
	}
	// if server.NewAgent == nil {
	// 	xlog.WarningLog(proxy.GetSecenName(),"NewAgent must not be nil")
	// }

	server.ln = ln
	server.conns = make(ConnSet)

	msgParser = NewMsgParser(server.appConf.Msglen_size,server.appConf.Max_msglen)
}

func (server *TCPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()
	var tempDelay time.Duration
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			//临时错误才继续，其他错误就关闭监听
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				xlog.ErrorLog(proxy.GetSecenName(),"accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			xlog.ErrorLog(proxy.GetSecenName(),"TCPServer Accept erro:%v", err)
			return
		}
		xlog.DebugLog(proxy.GetSecenName(),"Accept other erro ", conn.RemoteAddr())
		tempDelay = 0
		// 添加连接
	    if !server.addConn(conn) {
	      continue
	    }
		
		// go func() {
		// 	agent.Run()
		//
		// 	// cleanup
		// 	tcpConn.Close()
		// 	server.mutexConns.Lock()
		// 	delete(server.conns, conn)
		// 	server.mutexConns.Unlock()
		// 	agent.OnClose()
		//
		// 	server.wgConns.Done()
		// }()
	}
}

//添加链接信息
func (server *TCPServer) addConn(conn net.Conn) bool  {
	server.mutexConns.Lock()
	if server.connetsize  >= server.appConf.Max_connect{
		server.mutexConns.Unlock()
		conn.Close()
		xlog.WarningLog(proxy.GetSecenName(),"超过最大链接数,当前连接数%d",server.connetsize)
		return  false
	}
	server.conns[conn] = struct{}{}
	server.connetsize = len(server.conns)
	server.mutexConns.Unlock()
	server.wgConns.Add(1)
	//创建封装的连接
	tcpConn := newTCPConn(conn,server.netObserver,server.appConf)
	//连接接收消息
	go server.ReceiveData(tcpConn)
	xlog.DebugLog(proxy.GetSecenName(),"当前连接数%d",server.GetConnectSize())
	return true
}

// 连接中读取数据
func (server *TCPServer) ReceiveData(tcpConn *TCPConn) {
	defer server.wgConns.Done() //关闭连接要释放
	for {
		err := tcpConn.ReadMsg()
		if err != nil { // 这里读到错误消息,关闭
			xlog.WarningLog(proxy.GetSecenName(), "read message err: %s ", err.Error())
			break // 关闭连接
		}
	}
	
	// 处理远端关闭
	server.RemoveConn(tcpConn)
}
//断开连接
func (server *TCPServer) RemoveConn(tcpConn *TCPConn)   {
	tcpConn.Close()                      // 关闭连接
	server.mutexConns.Lock()
	delete(server.conns, tcpConn.conn)
	server.connetsize = len(server.conns)
	server.mutexConns.Unlock()
	xlog.DebugLog(proxy.GetSecenName(),"当前连接数%d",server.GetConnectSize())
}

//获取连接数
func (server *TCPServer) GetConnectSize() int   {
	server.mutexConns.RLock()
	defer server.mutexConns.RUnlock()
	return server.connetsize
}


func (server *TCPServer) Close() {
	server.ln.Close() // 关闭监听
	server.wgLn.Wait()

	server.mutexConns.Lock()
	for conn := range server.conns {
		conn.Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()
	server.wgConns.Wait()
	fmt.Println("TCPServer Close")
}
