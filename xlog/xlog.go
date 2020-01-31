//  创建时间: 2019/10/22
//  作者: zjy
//  功能介绍:
//  $ log日志,输出封装,这里要调整最好不予App有依赖
// 为了提高逻辑线程处理效率
//  多 gocorutine 传入日志信息 ,单 gocorutine 向文件写入
//  每日一个文件夹,每天两小时对应日志进程文件

package xlog

import (
	. "../model"
	"../timeutil"
	"../xutil"
	"fmt"
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
	Normal = 0 << iota
	DebugLvl
	WarningLvl
	ErrorLvl
)

var (
	xlog         *Xlog
	isOutStd     bool // 是否在标准输出输入
	loglvlStrMap map[int16]string
	showLvl      int16 // 显示日志等级
)

type Xlog struct {
	baseLog    *log.Logger    //内置log库的处理
	logBufchan chan *LogModel
	info       *LogInitModel  //初始化
	writeFile  *os.File
	closelog   chan int
}

//

// 创建日志对象
func NewXlog(info *LogInitModel) *Xlog {
	xlog := new(Xlog)
	xlog.baseLog = log.New(os.Stdout, "", log.LstdFlags)
	xlog.logBufchan = make(chan *LogModel, info.LogQueueCap)
	xlog.closelog = make(chan int)
	xlog.info = info
	initXlog()
	return xlog
}

func initXlog() {
	loglvlStrMap[Normal] = ""
	loglvlStrMap[DebugLvl] = "调试"
	loglvlStrMap[WarningLvl] = "警告"
	loglvlStrMap[ErrorLvl] = "错误"
}

//设置日志等级并设置是否在控制台显示
func SetShowLog(outStd bool,showlel  int16)  {
	isOutStd = outStd
	showLvl = showlel
}
func DebugLog(scenename string,format string,v ...interface{}) {
	addLogToLogBufchan(DebugLvl,scenename,format,v...)
}
func WarningLog(scenename string,format string, v ...interface{}) {
	addLogToLogBufchan(WarningLvl,scenename,format,v...)
}

func ErrorLog(scenename string,format string, v ...interface{}) {
	addLogToLogBufchan(ErrorLvl,scenename,format,v...)
}
// 向log日志队列中写日志信息
func addLogToLogBufchan(loglvl int16,scenename string,format string, v ...interface{}) {
	lm  := new(LogModel)
	lm.OutStr = fmt.Sprintf(format, v...)
	lm.LogGenerateTime = timeutil.GetCurrentTimeNano()
	lm.LogLvel = loglvl
	lm.SceneName = scenename
	xlog.logBufchan <- lm
}

func (xl *Xlog) writeLogToFile(lm *LogModel) {
	xl.NewLogDir(lm.LogGenerateTime)  //查看目录是否存在
	xl.NewLogFile(lm.LogGenerateTime,lm.SceneName) //创建文件
	xl.setOutFile()
	xl.setOutPrefix(lm.LogLvel)
	xl.baseLog.Println(lm.OutStr) //向输出流输出字符串
	xl.writeFile.Close() // 最后关闭文件
}

func (xl *Xlog) setOutFile() {
	//如果不能输出日志都在标准中输出
	if xl.writeFile == nil || !CanLogBylvl(DebugLvl)  {
		xl.baseLog.SetOutput(os.Stdout)
	} else if isOutStd && xl.writeFile != nil {
		xl.baseLog.SetOutput(io.MultiWriter(os.Stdout, xl.writeFile))
	} else {
		xl.baseLog.SetOutput(xl.writeFile)
	}
}
// 创建日志日期路径
func (xl *Xlog) NewLogDir(currentNano int64) {
	dirs := path.Join(xl.info.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano))
	xutil.MakeDir(dirs)
}

func (xl *Xlog) NewLogFile(currentNano int64,scenename string) {
	filename := timeutil.GetYearMonthDayHourFromatStr(currentNano) + xl.info.ServerName + "_" + scenename+ ".log"
	str := path.Join(xl.info.LogsPath, timeutil.GetYearMonthDayFromatStr(currentNano), filename)
	fmt.Println(str)
	tempfile, err := os.OpenFile(str, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		xlog.baseLog.Println(err)
	}
	xl.writeFile = tempfile
}

func (xl *Xlog) setOutPrefix(reqlvl int16) {
	if prefixStr, ok := loglvlStrMap[reqlvl]; ok {
		xl.baseLog.SetPrefix(prefixStr)
	}
}

//日志执行逻辑线程
func (xl *Xlog) run() {
	
	//拉起错误
	defer func() {
		if err := recover(); err != nil {
			ErrorLogInterfaceParam(err)
		}
	}()
	
	for {
		select {
		case logmodel := <-xl.logBufchan: // 获取队列数据
			xl.writeLogToFile(logmodel)
		case listenCloseCode := <-xl.closelog:
			xl.baseLog.Println("关闭日志码 = ", listenCloseCode)
			break // 这里就跳出循环
		}
	}
	// 结束日志线
	xl.close()
}


func ErrorLogInterfaceParam(logstr interface{}) {
	if logstrtem, ok := logstr.(string); ok {
		ErrorLog("Normol",logstrtem)
	}
}


//如果不能输出日志都在标准中输出
func CanLogBylvl(loglvl int16) bool {
	return (loglvl & showLvl) != 0
}

// 关闭日志
func CloseLog(colseCode int) {
	xlog.closelog <- colseCode
}

func (xl *Xlog) close() {
	close(xl.logBufchan)
	close(xl.closelog)
	if xl.writeFile != nil {
		xl.writeFile.Close()
	}
	for k, _ := range loglvlStrMap {
		delete(loglvlStrMap, k)
	}
}
