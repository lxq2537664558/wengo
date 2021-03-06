/*
创建时间: 2019/12/22
作者: zjy
功能介绍:
原子开关用于检测
*/

package model

import "sync/atomic"

type AtomicInt32FlagModel struct {
	 checkFlag int32
}

const
(
	CloseFlag  = 0
	OpenFlag = 1
)

func NewAtomicInt32Flag() *AtomicInt32FlagModel {
	return new(AtomicInt32FlagModel)
}

// 检测是否关闭
func (af *AtomicInt32FlagModel) IsClosed() bool {
	return atomic.LoadInt32(&af.checkFlag) == CloseFlag
}

// 检测是否开启
func (af *AtomicInt32FlagModel) IsOpen() bool {
	return atomic.LoadInt32(&af.checkFlag) == OpenFlag
}

func (af *AtomicInt32FlagModel) Close() {
	af.SetFlag(CloseFlag)
}

func (af *AtomicInt32FlagModel) Open() {
	af.SetFlag(OpenFlag)
}

func (af *AtomicInt32FlagModel) SetFlag(flagval int32) {
	atomic.StoreInt32(&af.checkFlag,flagval)
}

