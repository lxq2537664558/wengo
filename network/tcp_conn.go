// 1.封装tcp连接,改结构体只负责写的工作，
// 2.读的工作 远端关闭根据实现的是服务器或者客户端不同

package network

import (
	"errors"
	"github.com/showgo/csvdata"
	"github.com/showgo/proxy"
	"github.com/showgo/timeutil"
	"github.com/showgo/xlog"
	"net"
	"sync"
)

type ConnSet map[net.Conn]struct{}

type TCPConn struct {
	conn        net.Conn
	netObserver NetWorkObserver
	wgConns     *sync.WaitGroup
	sync.RWMutex                            // 主要作用,防止向关闭后的通道中写入数据
	closeFlag   bool                        // 检测关闭标志 这里不用锁使用原子数据更效率
	writeChan   chan []byte                 // 写的通道，我服务器写的消息先写入通道再用连接传出去
	RecMutex    sync.RWMutex                // 接收锁
	lastRecTime int64                       // 最後一次收包时间
	RecMsgPs    int                         // 每秒收包个数
	maxRecMsgPs int                         // 每秒最大能收多少个
	appConf   *csvdata.Appconf
}


func init() {

}

func newTCPConn(conn net.Conn, netOb NetWorkObserver,appconf   *csvdata.Appconf) *TCPConn {
	tcpConn := new(TCPConn)
	tcpConn.conn = conn
	tcpConn.netObserver = netOb
	tcpConn.appConf = appconf
	tcpConn.writeChan = make(chan []byte, tcpConn.appConf.Write_cap_num)
	// tcpConn.msgParser = msgParser
	go tcpConn.writeChanData()                    // 写协程
	tcpConn.netObserver.OnNetWorkConnect(tcpConn) // 通知其他模块已经连接
	
	return tcpConn
}

// 取通道的数据给连接
func (tcpConn *TCPConn) writeChanData() {
	// 这里接收写的通道，没有数据会一直阻塞，直到通道关闭
	for b := range tcpConn.writeChan {
		if b == nil {
			break
		}
		_, err := tcpConn.conn.Write(b)
		if err != nil {
			break
		}
	}

	// 这里是因为写的通道满了关闭连接,主动关闭连接,当读协程读取关闭时
	tcpConn.conn.Close()
	// 已经关闭
	tcpConn.Lock()
	tcpConn.closeFlag  = true
	tcpConn.Unlock()
}

func (tcpConn *TCPConn) doDestroy() {
	tcpConn.conn.(*net.TCPConn).SetLinger(0)
	tcpConn.conn.Close()
	
	if !tcpConn.closeFlag {
		tcpConn.closeFlag = true
		close(tcpConn.writeChan)
	}
}

func (tcpConn *TCPConn) Destroy() {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	tcpConn.doDestroy()
}

func (tcpConn *TCPConn) Close() {
	xlog.DebugLog(proxy.GetSecenName(), "TCPConn Close")
	
	tcpConn.Lock()
	// 已经关闭
	if tcpConn.closeFlag {
		tcpConn.Unlock()
		xlog.DebugLog(proxy.GetSecenName(), "当前连接已经关闭")
		return
	}
	tcpConn.closeFlag = true
	tcpConn.doWrite(nil)
	tcpConn.Unlock()
	xlog.DebugLog(proxy.GetSecenName(), "closeFlag.Close1", tcpConn.closeFlag)
	// 通知其他模块
	tcpConn.netObserver.OnNetWorkClose(tcpConn)
}

func (tcpConn *TCPConn) doWrite(b []byte) {
	// 写的队列被撑满时
	if len(tcpConn.writeChan) == cap(tcpConn.writeChan) {
		xlog.DebugLog(proxy.GetSecenName(), "close conn: channel full")
		tcpConn.doDestroy() // 这里要主动断开避免阻塞 当前调用协程
		return
	}
	xlog.DebugLog(proxy.GetSecenName(), "TCPConn doWrite", b)
	tcpConn.writeChan <- b
}

// b must not be modified by the others goroutines
func (tcpConn *TCPConn) Write(b []byte) {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	// 已经关闭
	if tcpConn.closeFlag || b == nil {
		xlog.DebugLog(proxy.GetSecenName(), "当前连接状态", tcpConn.closeFlag, "写入管道数据", b)
		return
	}
	tcpConn.doWrite(b)
}

func (tcpConn *TCPConn) Read(b []byte) (int, error) {
	return tcpConn.conn.Read(b)
}

func (tcpConn *TCPConn) LocalAddr() net.Addr {
	return tcpConn.conn.LocalAddr()
}

func (tcpConn *TCPConn) RemoteAddr() net.Addr {
	return tcpConn.conn.RemoteAddr()
}

// 用解析对象读
func (tcpConn *TCPConn) ReadMsg() (error) {
	data, err := msgParser.Read(tcpConn)
	if err != nil {
		return err
	}
	// 这里应该进入队列
	maincmd, subcmd, msgdat, erro := msgParser.UnpackOne(data)
	if erro != nil {
		return erro
	}
	xlog.DebugLog(proxy.GetSecenName(), "ReadMsg", maincmd, subcmd, string(msgdat))

	msgData := NewMsgData(tcpConn,maincmd,subcmd,msgdat)
	if msgData == nil {
		return errors.New("NewMsgData is nil")
	}
	tcpConn.netObserver.OnNetWorkRead(msgData)
	return nil
}

// 写单个消息
func (tcpConn *TCPConn) WriteOneMsg(maincmd, subcmd uint16, msg []byte) error {
	data, erro := msgParser.PackOne(maincmd, subcmd, msg)
	if erro != nil {
		xlog.DebugLog(proxy.GetSecenName(), "WriteOneMsg erro : ", erro.Error())
		return erro
	}
	// 向写通道投递数据
	tcpConn.Write(data)
	return nil
}

// 将消息体构建为[]byte数组，最终要发出去的单包
func (tcpConn *TCPConn) GetOneMsgByteArr(maincmd, subcmd uint16, msg []byte) ([]byte, error) {
	return msgParser.PackOne(maincmd, subcmd, msg)
}

// 一起写多个数据包
// 每个包的数据 由GetOneMsgByteArr构建
// 但是不能超过 writeChan 通道的大小
func (tcpConn *TCPConn) WriteMsg(args ...[]byte) error {
	var msgLen uint32
	//计算消息长度
	for i := 0; i < len(args); i++ {
		msgLen += uint32(len(args[i]))
	}
	// check len
	if msgLen > msgParser.maxMsgLen {
		return errors.New("message too long")
	} else if msgLen < msgParser.minMsgLen {
		return errors.New("message too short")
	}
	//构建所有数据包
	msg := make([]byte, msgLen)
	l := 0
	for i := 0; i < len(args); i++ {
		copy(msg[l:], args[i])
		l += len(args[i])
	}
	tcpConn.Write(msg)
	return nil
}

//查看是否可以读取
func (tcpConn *TCPConn) checkCanRead() bool {
	tcpConn.RecMutex.Lock()
	defer tcpConn.RecMutex.Unlock()
	currentTime := timeutil.GetCurrentTimeS()
	//在同一秒内
	if tcpConn.lastRecTime == currentTime {
		//收包的数量超过每秒最大限制数量
		if 	tcpConn.RecMsgPs >= tcpConn.maxRecMsgPs {
			return false
		}
		tcpConn.RecMsgPs ++
		return true
	}
	//过了一秒重置变量
	tcpConn.RecMsgPs = 0
	tcpConn.lastRecTime = currentTime
	return true
}
