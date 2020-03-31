//  创建时间: 2019/10/22
//  作者: zjy
//  功能介绍:
//  $ log日志,输出封装,这里要调整最好不予App有依赖
// 为了提高逻辑线程处理效率
//  多 gocorutine 传入日志信息 ,单 gocorutine 向文件写入
//  每日一个文件夹,每天两小时对应日志进程文件

package xlog

import (
	"fmt"
	"github.com/showgo/model"
	"github.com/showgo/timeutil"
	"github.com/showgo/xutil"
	"io"
	"log"
	"os"
	"path"
	"sync"
)

const (
	CloseType_nomarl = 1
)

// 日志等级定义
const (
	Normal     = 0
	DebugLvl   = 1 << 0
	WarningLvl = 1 << 1
	ErrorLvl   = 1 << 2
)

var (
	_xlog        *Xlog // 日志执行者对象
	loglvlStrMap map[uint16]string
	_wg          sync.WaitGroup // 为保证程序统一退出这里加个等待
)

type Xlog struct {
	initInfo     *LogInitModel  // 初始化
	baseLog      *log.Logger    // 内置log库的处理
	writeFile    *os.File       // 写日志的文件对象
	logBufchan   chan *LogModel // 日志信息
	closeFlag    *model.AtomicInt32FlagModel
	logmodelPool sync.Pool
}

// 创建日志对象
func NewXlog(info *LogInitModel) bool {
	_xlog = new(Xlog)
	if _xlog == nil {
		fmt.Println("_NewXlog xlog is nil")
		return false
	}
	_xlog.baseLog = log.New(os.Stdout, "", 0)
	_xlog.logBufchan = make(chan *LogModel, info.Volatile.LogQueueCap)
	_xlog.initInfo = info
	
	loglvlStrMap = make(map[uint16]string)
	initXlog()
	_wg.Add(1)
	go _xlog.run()
	return true
}

func initXlog() {
	
	loglvlStrMap[Normal] = "无"
	loglvlStrMap[DebugLvl] = "调试|"
	loglvlStrMap[WarningLvl] = "警告|"
	loglvlStrMap[ErrorLvl] = "错误|"
	_xlog.closeFlag = model.NewAtomicInt32Flag()
	_xlog.closeFlag.Open()
	_xlog.logmodelPool.New = func() interface{} {
		return new(LogModel)
	}
}

// 设置日志等级并设置是否在控制台显示 目前这两个经常改变
func SetShowLogAndStartLog(restmodel VolatileLogModel) bool {
	if _xlog == nil {
		fmt.Println("SetShowLogAndStartLog xlog is nil")
		return false
	}
	_xlog.initInfo.Volatile = restmodel
	return true
}

func DebugLog(scenename string, format string, v ...interface{}) {
	addLogToLogBufchan(DebugLvl, scenename, format, v...)
}
func WarningLog(scenename string, format string, v ...interface{}) {
	addLogToLogBufchan(WarningLvl, scenename, format, v...)
}

func ErrorLog(scenename string, format string, v ...interface{}) {
	addLogToLogBufchan(ErrorLvl, scenename, format, v...)
}

// 向log日志队列中写日志信息
func addLogToLogBufchan(loglvl uint16, scenename string, format string, v ...interface{}) {
	if _xlog == nil {
		fmt.Println("addLogToLogBufchan xlog is nil")
		return
	}
	if _xlog.closeFlag.IsClosed() {
		return
	}
	// 未设置对应的日志等级就不能打印
	if !canLogBylvl(loglvl) {
		return
	}
	tem := _xlog.logmodelPool.Get() // TODO 这里可以优化日志对象创建
	if  tem == nil {
		return
	}
	lm,ok := tem.(*LogModel)
	if !ok {
		return
	}
	lm.OutStr = fmt.Sprintf(format, v...)
	lm.LogGenerateTime = timeutil.GetCurrentTimeNano()
	lm.LogLvel = loglvl
	lm.SceneName = scenename
	_xlog.logBufchan <- lm
}

func (xl *Xlog) writeLogToFile(lm *LogModel) {
	if lm == nil {
		return
	}
	isOk := xl.newLogsDir(lm.LogGenerateTime) // 查看目录是否存在
	if !isOk {
		return
	}
	isOk = xl.newLogFile(lm.LogGenerateTime, lm.SceneName) // 创建文件
	if !isOk {
		return
	}
	xl.setOutFile()
	xl.setOutPrefix(lm.LogLvel, lm.LogGenerateTime)
	xl.baseLog.Println(lm.OutStr) // 向输出流输出字符串
	_xlog.logmodelPool.Put(lm)    // 放回池子
	xl.writeFile.Close()          // 最后关闭文件
}

// 创建日志日期路径
func (xl *Xlog) newLogsDir(currentNano int64) bool {
	dirs := path.Join(xl.initInfo.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano))
	return xutil.MakeDirAll(dirs)
}

func (xl *Xlog) newLogFile(currentNano int64, scenename string) bool {
	// 每两个小时一个文件
	filename := timeutil.GetYearMonthDayHourFromatStrBySpan(currentNano, xl.initInfo.Volatile.FileTimeSpan) + "_" + xl.initInfo.ServerName + "_" + scenename + ".log"
	str := path.Join(xl.initInfo.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano), filename)
	tempfile, err := os.OpenFile(str, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("打开日志文件错误 = ", err)
		tempfile.Close()
		return false
	}
	xl.writeFile = tempfile
	return true
}

func (xl *Xlog) setOutFile() {
	if xl.writeFile == nil {
		xl.baseLog.SetOutput(os.Stdout)
	} else if xl.initInfo.Volatile.IsOutStd && xl.writeFile != nil {
		xl.baseLog.SetOutput(io.MultiWriter(os.Stdout, xl.writeFile))
	} else {
		xl.baseLog.SetOutput(xl.writeFile)
	}
}

func (xl *Xlog) setOutPrefix(reqlvl uint16, currentNano int64) {
	// 清除日志时间
	if prefixStr, ok := loglvlStrMap[reqlvl]; ok {
		// 日志等级与生成时间
		xl.baseLog.SetPrefix(fmt.Sprintf("%s%s\t", prefixStr, timeutil.GetTimeALLStr(currentNano)))
	} else {
		xl.baseLog.SetPrefix(timeutil.GetTimeALLStr(currentNano))
		
	}
}

// 日志执行逻辑线程
func (xl *Xlog) run() {
	
	defer _wg.Done()
	// 拉起宕机
	defer RecoverToStd()

ENDLOOP:
	for {
		select {
		case logmodel := <-xl.logBufchan: // 获取队列数据
			if logmodel == nil {
				break ENDLOOP // 这里就跳出循环
			}
			xl.writeLogToFile(logmodel)
		}
	}
}

// 如果不能输出日志都在标准中输出
func canLogBylvl(loglvl uint16) bool {
	if _xlog == nil {
		return false
	}
	return (loglvl & _xlog.initInfo.Volatile.ShowLvl) != 0
}

// 关闭日志
func CloseLog() {
	_xlog.closeFlag.Close()
	_xlog.logBufchan <- nil
	_xlog.onClose()
}

func (xl *Xlog) onClose() {
	_wg.Wait()
	close(xl.logBufchan)
	if xl.writeFile != nil {
		xl.writeFile.Close()
	}
	loglvlStrMap = nil
	fmt.Println("Log Close")
}
