/*
创建时间: 2020/2/3
作者: zjy
功能介绍:
事件系统
*/

package dispatch

import (
	"fmt"
	"github.com/showgo/model"
	"github.com/showgo/network"
	"github.com/showgo/proxy"
	"github.com/showgo/xcontainer/queue"
	"github.com/showgo/xlog"
	"sync"
)

var DispSys *DispatchSys

type DispatchNoticeFun func(interface{}) error // 对应的解析函数

type DispatchSys struct {
	eventmtx          sync.Mutex
	eventcond         *sync.Cond        //控制条件
	eventQueue        *queue.SafeQueue // 主逻辑事件队列
	
	netObserver       network.NetWorkObserver
	endFlag           *model.AtomicInt32FlagModel
	wg                sync.WaitGroup
	disPatchParseFuns []DispatchNoticeFun
	dispatchDatapool sync.Pool
}

// go自动调用 初始化管理变量
func init() {
	// DispSys = NewDispatchSys()
}

func NewDispatchSys() *DispatchSys {
	disp := new(DispatchSys)
	disp.disPatchParseFuns = make([]DispatchNoticeFun, DisPatch_max)
	disp.eventQueue = queue.NewSafeQueue()
	disp.eventcond = sync.NewCond(&disp.eventmtx) //初始化条件变量
	disp.Init()
	disp.endFlag = model.NewAtomicInt32Flag()
	disp.endFlag.Open()    // 打开标志
	disp.dispatchDatapool.New = DisPatchDataPoolNewFun  //设置创建对象函数
	disp.wg.Add(1)
	go disp.OnQueueEvent() // 启动
	return disp
}

func (this *DispatchSys) Init() {
	this.RegisterDispatchEvent(Timer_Event,nil)
	this.RegisterDispatchEvent(NetWorkAccept_Event,this.NoticeNetWorkAccept)
	this.RegisterDispatchEvent(NetWorkRead_Event,this.NoticeNetWorkRead)
	this.RegisterDispatchEvent(NetWorkClose_Event,this.NoticeNetWorkClose)
}

func (this *DispatchSys) RegisterDispatchEvent(event int,noticeFun DispatchNoticeFun)  {
	this.disPatchParseFuns[event] = noticeFun
}
// 设置消息处理观察者
func (this *DispatchSys) SetObserver(netobserver network.NetWorkObserver) {
	this.netObserver = netobserver
}

// 投递定时器事件
func (this *DispatchSys) PostTimer(timerParam interface{}) error {
	data := this.GetDisPatchDataByPool(Timer_Event, timerParam)
	this.eventQueue.PushBack(data)
	this.eventcond.Signal()
	return nil
}

// 网络连接事件
func (this *DispatchSys) OnNetWorkConnect(conn network.Conner) error {
	data := this.GetDisPatchDataByPool(NetWorkAccept_Event, conn)
	this.eventQueue.PushBack(data)
	this.eventcond.Signal()
	return nil
}

// 网络读取事件
func (this *DispatchSys) OnNetWorkRead(msgdata *network.MsgData) error {
	data := this.GetDisPatchDataByPool(NetWorkRead_Event, msgdata)
	this.eventQueue.PushBack(data)
	this.eventcond.Signal()
	return nil
}

// 网络关闭事件
func (this *DispatchSys) OnNetWorkClose(conn network.Conner) error {
	data := this.GetDisPatchDataByPool(NetWorkClose_Event, conn)
	this.eventQueue.PushBack(data)
	this.eventcond.Signal()
	return nil
}

func (this *DispatchSys) GetDisPatchDataByPool(dtype int, val interface{}) *DisPatchData {
	dpd := this.dispatchDatapool.Get()
	if dpd == nil {
		xlog.WarningLog(proxy.GetSecenName(), " dispatchDatapool is nil")
		return  NewDisPatchData(dtype,val)
	}
	data,ok:= dpd.(*DisPatchData)
	if !ok {
		xlog.WarningLog(proxy.GetSecenName(), " GetDisPatchDataByPool is nil")
		return  NewDisPatchData(dtype,val)
	}
	data.SetDisPatchData(dtype,val)
	return data
}

func (this *DispatchSys) OnQueueEvent() {
	defer xlog.RecoverToLog()	//拉起错误避免宕机
	defer this.wg.Done()
	for {
		if this.endFlag.IsClosed() {
			xlog.DebugLog(proxy.GetSecenName(), " OnQueueEvent end Run")
			break
		}
		this.eventmtx.Lock()
		//当没有事件发生时，要阻塞
		if this.eventQueue.Len() == 0 {
			this.eventcond.Wait()
		}
		this.eventmtx.Unlock()
		// 取队列数据
		dipatch, erro := this.eventQueue.PopFront()
		if erro != nil {
			this.dispatchDatapool.Put(dipatch) //放回到池子
			xlog.DebugLog(proxy.GetSecenName(), " OnQueueEvent 取出数据错误", erro.Error())
			continue
		}
		// 解析数据
		data, ok := dipatch.(*DisPatchData)
		if !ok {
			xlog.DebugLog(proxy.GetSecenName(), " OnQueueEvent 解析数据失败")
			continue
		}
		if data.dipatchType < 0 || data.dipatchType >= DisPatch_max {
			continue
		}
		// 查找对应的方法处理数据
		paseFun := this.disPatchParseFuns[data.dipatchType]
		if paseFun != nil {
			erro := paseFun(data.val)
			if erro != nil {
				xlog.DebugLog(proxy.GetSecenName(), " paseFun dipatchType = %d, err %s", data.dipatchType, erro.Error())
			}
		}
		this.dispatchDatapool.Put(dipatch) //放回到池子
		xlog.DebugLog(proxy.GetSecenName(), " this.eventQueue剩余需要处理的数据 = %d", this.eventQueue.Len())
	}
}

// 关闭系统
func (this *DispatchSys) Release() {
	this.eventcond.Signal() //唤醒协程
	this.endFlag.Close()
	xlog.DebugLog(proxy.GetSecenName(), "还有未处理的事件%d",this.eventQueue.Len())
	this.eventQueue.Clear()
	this.wg.Wait() // 等待所有的协程回收
	fmt.Println("DispatchSys Release")
}
