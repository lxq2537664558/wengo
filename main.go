// 创建时间: 2019-10-2019/10/17
// 作者: zjy
// 功能介绍:
// 1.主要入口
// 2.
// 3.
package main

import (
	"fmt"
	"github.com/showgo/model"
	"github.com/showgo/xlog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// main 初始化工作
func init() {
}

// 各服务器主入口
func main() {
	
	logInit := model.LogInitModel{
		"Login",
		"./logs",
		model.RestLogModel{
			10,
			true,
			7,
		},
	}
	
	isOK :=  xlog.NewXlog(logInit)
	if !isOK {
		return
	}
	// 设置最大运行核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 10 ; i++ {
		go func(a int) {
			xlog.WarningLog("main ","gocorutine %d ",a)
			xlog.DebugLog("main ","gocorutine %d",a)
			xlog.ErrorLog("main ","gocorutine %d",a)
		}(i)
	}
	
	time.Sleep(time.Second * 10)
	xlog.CloseLog(xlog.CloseType_nomarl)
	time.Sleep(time.Second * 10)
	// app.GetStart()
	// 等待退出 在app 退出后整个程序退出
	fmt.Println("Main Exit")
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGKILL)
	select {
	case s := <-c:
		// 收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
		break
	}
}

