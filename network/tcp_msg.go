package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"github.com/showgo/xutil"
	"math"
)

// 包格式
// --------------
// | 包总长 | 主命令 | 字命令 | datalen | data |
// --------------
type MsgParser struct {
	msgLenSize uint8            // 包长度字节大小
	msgheadLen uint32           // 包头默认需要最小字节数
	minMsgLen  uint32           // 最小长度
	maxMsgLen  uint32           // 最大长度
}

var msgParser    *MsgParser  //数据包解析对象

var	byteOrder  binary.ByteOrder // 设置网址字节序

func NewMsgParser(msgLenSize uint8, maxMsgLen uint32) *MsgParser {
	p := new(MsgParser)
	p.msgLenSize = msgLenSize
	var head TcpMsgHead
	p.msgheadLen = uint32(binary.Size(head)) //包头需要长度
	p.setMsgLen(maxMsgLen)
	// 设置大段序列与小端序列
	if xutil.IsLittleEndian() {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}
	return p
}

// 设置消息长度
func (p *MsgParser) setMsgLen(maxMsgLen uint32) {
	//数据包最小长度 | 包总长大小 | 主命令大小 | 字命令大小 | datalen大小 |
	p.minMsgLen = uint32(p.msgLenSize) + p.msgheadLen
	if maxMsgLen != 0 {
		p.maxMsgLen = maxMsgLen
	}
	var max uint32
	switch p.msgLenSize {
	case 2:
		max = math.MaxUint16
	case 4:
		max = math.MaxUint32
	}
	if p.minMsgLen > max { //不能超过设置的
		p.minMsgLen = max
	}
	if p.maxMsgLen > max {
		p.maxMsgLen = max
	}
}

// goroutine safe
func (p *MsgParser) Read(conn *TCPConn) ([]byte, error) {
	//根据长度字节大小解析第一个长度
	bufMsgLen := make([]byte,p.msgLenSize)
	// read len
	n, err := conn.Read(bufMsgLen);
	if err != nil {
		xlog.DebugLog(proxy.GetSecenName(), "read msglen erro : %s", err.Error())
		return nil, err
	}
	// parse len
	msgLen := p.byteArrToMsgLen(bufMsgLen)
	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}
	// data 读取
	// 已经提取了消息长度的字节 剩余| 主命令 | 字命令 | datalen | data | 这里扣除长度的字节
	datalen := msgLen - uint32(p.msgLenSize)
	msgData := make([]byte, datalen) //TODO 可以优化接收消息的buf
	n, err = conn.Read(msgData);
	if err != nil {
		return nil, err
	}
	//长度与数据不符
	if  datalen != uint32(n)  {
		return nil, errors.New(fmt.Sprintf("msglen = %d, readlen = %d",datalen,n))
	}
	return msgData, nil
}

//根据字节数组解析长度
func (p *MsgParser) byteArrToMsgLen(bArr []byte) uint32 {
	switch p.msgLenSize {
	case 2: // uint16 解析
		return uint32(ByteArrToUint16(bArr))
	case 4: // uint32 解析
		return ByteArrToUint32(bArr)
	}
	xlog.ErrorLog(proxy.GetSecenName(),"not set msg size ")
	return 0 // 无效
}

func ByteArrToUint32(bArr []byte) uint32 {
	return byteOrder.Uint32(bArr)
}

func ByteArrToUint16(bArr []byte) uint16 {
	return byteOrder.Uint16(bArr)
}

// param | 主命令 | 字命令 | datalen | data |
func (p *MsgParser)UnpackOne(readb []byte) (maincmd uint16, subcmd uint16, msg []byte, err error) {
	reader := bytes.NewBuffer(readb)
	var head TcpMsgHead
	err = binary.Read(reader, byteOrder, &head)
	if  err != nil {
		return
	}
	maincmd = head.MainCmd
	subcmd = head.SubCmd
	msg = make([]byte, head.Datalen)
	err = binary.Read(reader, byteOrder, &msg)
	return
}


// 打单包
// 使用 bytes.buffer 作为io.Writer中介
func (p *MsgParser)PackOne(maincmd, subcmd uint16, msg []byte) ( []byte, error) {
	datalen := uint32(len(msg))
	head := TcpMsgHead{
		MainCmd:    maincmd,
		SubCmd:     subcmd,
		Datalen:    datalen,
	}
	msgLen := p.minMsgLen + datalen
	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil,errors.New("message too short")
	}
	var resmsg []byte
	writer := bytes.NewBuffer(resmsg)
	//写长度
	switch p.msgLenSize {
	case 2: // uint16
		binary.Write(writer,byteOrder,uint16(msgLen))
	case 4: // uint32
		binary.Write(writer,byteOrder,uint32(msgLen))
	}
	//写头
	erro := binary.Write(writer, byteOrder, head)
	//写数据
	erro = binary.Write(writer,byteOrder,msg)
	return writer.Bytes(),erro
}
