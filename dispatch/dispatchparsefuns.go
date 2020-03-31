/*
创建时间: 2020/2/29
作者: zjy
功能介绍:

*/

package dispatch

import (
	"errors"
	"github.com/showgo/network"
)

// 网络连接事件
func (this *DispatchSys) NoticeNetWorkAccept(val interface{}) error {
	conn,ok:= val.(network.Conner)
	if !ok {
		return errors.New("NoticeNetWorkAccept  Assert network.Conner type Erro")
	}
	if this.netObserver == nil {
		return  errors.New("NoticeNetWorkAccept  netObserver is nil")
	}
	return this.netObserver.OnNetWorkConnect(conn)
}

// 网络读取事件
func (this *DispatchSys) NoticeNetWorkRead(val interface{}) error {
	msgData,ok:= val.(*network.MsgData)
	if !ok {
		return errors.New("NoticeNetWorkRead  Assert network.MsgData type Erro")
	}
	if this.netObserver == nil {
		return  errors.New("NoticeNetWorkRead  netObserver is nil")
	}
	return this.netObserver.OnNetWorkRead(msgData)
}

// 网络关闭事件
func (this *DispatchSys) NoticeNetWorkClose(val interface{}) error {
	conn,ok:= val.(network.Conner)
	if !ok {
		return errors.New("NoticeNetWorkClose  Assert network.Conner type Erro")
	}
	if this.netObserver == nil {
		return  errors.New("NoticeNetWorkClose  netObserver is nil")
	}
	return this.netObserver.OnNetWorkClose(conn)
}
