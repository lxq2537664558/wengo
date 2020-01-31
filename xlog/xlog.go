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
	. "github.com/showgo/model"
	"github.com/showgo/timeutil"
	"github.com/showgo/xutil"
	"io"
	"log"
	"os"
	"path"
)

const (
	CloseType_nomarl = iota
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
	loglvlStrMap map[int16]string
)

type Xlog struct {
	initInfo   LogInitModel // 初始化
	baseLog    *log.Logger  // 内置log库的处理
	writeFile  *os.File     // 写日志的文件对象
	logBufchan chan *LogModel
	closelog   chan int
}

// 创建日志对象
func NewXlog(info LogInitModel) bool {
	_xlog = new(Xlog)
	if _xlog == nil {
		fmt.Errorf("_NewXlog xlog is nil")
		return false
	}
	_xlog.baseLog = log.New(os.Stdout, "", log.LstdFlags)
	_xlog.logBufchan = make(chan *LogModel, info.LogQueueCap)
	_xlog.closelog = make(chan int)
	_xlog.initInfo = info
	loglvlStrMap = make(map[int16]string)
	initXlog()
	go _xlog.run()
	return true
}

func initXlog() {
	loglvlStrMap[Normal] = "无"
	loglvlStrMap[DebugLvl] = "调试||"
	loglvlStrMap[WarningLvl] = "警告||"
	loglvlStrMap[ErrorLvl] = "错误||"
}

// 设置日志等级并设置是否在控制台显示 目前这两个经常改变
func SetShowLogAndStartLog(restmodel RestLogModel) bool {
	if _xlog == nil {
		fmt.Errorf("SetShowLogAndStartLog xlog is nil")
		return false
	}
	_xlog.initInfo.LogQueueCap = restmodel.LogQueueCap
	_xlog.initInfo.IsOutStd = restmodel.IsOutStd
	_xlog.initInfo.ShowLvl = restmodel.ShowLvl
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
func addLogToLogBufchan(loglvl int16, scenename string, format string, v ...interface{}) {
	// 未设置对应的日志等级就不能打印
	if !canLogBylvl(loglvl) {
		return
	}
	lm := new(LogModel)
	lm.OutStr = fmt.Sprintf(format, v...)
	lm.LogGenerateTime = timeutil.GetCurrentTimeNano()
	lm.LogLvel = loglvl
	lm.SceneName = scenename
	_xlog.logBufchan <- lm
}

func (xl *Xlog) writeLogToFile(lm *LogModel) {
	isOk := xl.NewLogDir(lm.LogGenerateTime) // 查看目录是否存在
	if !isOk {
		return
	}
	isOk = xl.NewLogFile(lm.LogGenerateTime, lm.SceneName) // 创建文件
	if !isOk {
		return
	}
	xl.setOutFile()
	xl.setOutPrefix(lm.LogLvel)
	xl.baseLog.Println(lm.OutStr) // 向输出流输出字符串
	xl.writeFile.Close()          // 最后关闭文件
}

func (xl *Xlog) setOutFile() {
	if xl.writeFile == nil {
		xl.baseLog.SetOutput(os.Stdout)
	} else if _xlog.initInfo.IsOutStd && xl.writeFile != nil {
		xl.baseLog.SetOutput(io.MultiWriter(os.Stdout, xl.writeFile))
	} else {
		xl.baseLog.SetOutput(xl.writeFile)
	}
}

// 创建日志日期路径
func (xl *Xlog) NewLogDir(currentNano int64) bool {
	dirs := path.Join(xl.initInfo.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano))
	return xutil.MakeDir(dirs)
}

func (xl *Xlog) NewLogFile(currentNano int64, scenename string) bool {
	filename := timeutil.GetYearMonthDayHourFromatStr(currentNano) + "_" + xl.initInfo.ServerName + "_" + scenename + ".log"
	str := path.Join(xl.initInfo.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano), filename)
	tempfile, err := os.OpenFile(str, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		_xlog.baseLog.Println(err)
		return false
	}
	xl.writeFile = tempfile
	return true
}

func (xl *Xlog) setOutPrefix(reqlvl int16) {
	if prefixStr, ok := loglvlStrMap[reqlvl]; ok {
		xl.baseLog.SetPrefix(prefixStr)
	}
}

// 日志执行逻辑线程
func (xl *Xlog) run() {
	
	// 拉起宕机
	defer func() {
		if err := recover(); err != nil {
			ErrorLogInterfaceParam(err)
		}
	}()

ENDLOOP:
	for {
		select {
		case logmodel := <-xl.logBufchan: // 获取队列数据
			xl.writeLogToFile(logmodel)
		case listenCloseCode := <-xl.closelog:
			fmt.Println("日志关闭码 = ", listenCloseCode)
			break ENDLOOP // 这里就跳出循环
		}
	}
	xl.close()
}

func ErrorLogInterfaceParam(logstr interface{}) {
	if logstrtem, ok := logstr.(string); ok {
		fmt.Errorf("Normol", logstrtem)
	}
}

// 如果不能输出日志都在标准中输出
func canLogBylvl(loglvl int16) bool {
	if _xlog == nil {
		return false
	}
	return (loglvl & _xlog.initInfo.ShowLvl) != 0
}

// 关闭日志
func CloseLog(colseCode int) {
	_xlog.closelog <- colseCode
}

func (xl *Xlog) close() {
	fmt.Println("Close log")
	close(xl.logBufchan)
	close(xl.closelog)
	if xl.writeFile != nil {
		xl.writeFile.Close()
	}
	for k, _ := range loglvlStrMap {
		delete(loglvlStrMap, k)
	}
}
