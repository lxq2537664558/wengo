// 创建时间: 2019-10-2019/10/17
// 作者: zjy
// 功能介绍:
// 1.主要入口
// 2.
// 3.
package main

import (
	"fmt"
	"./app"
	"os"
	"os/signal"
	 "path"
	"runtime"
	"syscall"
)

// main 初始化工作
func init() {
}

// 各服务器主入口
func main() {
	// Gwp.Add(1) // 控制主线程退出
	app.GetStart()
	// 设置最大运行核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	paths := path.Join("aa","bbb")
	fmt.Println(paths)
	app.App = app.NewServerApp()
	app.App.Start()
	
	// 等待退出 在app 退出后整个程序退出
	// Gwp.Wait()
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
