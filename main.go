// 创建时间: 2019-10-2019/10/17
// 作者: zjy
// 功能介绍:
// 1.主要入口
// 2.
// 3.
package main

import (
	"fmt"
	. "showgo/app"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)



// main 初始化工作
func init() {

	var a  = 0
	var kind  AppKind  = ItoAppKind(a)
	print("mianinit = ",kind.String())

}


//各服务器主入口
func main() {
		
	    // 设置最大运行核数
		runtime.GOMAXPROCS(runtime.NumCPU())
	   //系统监听
		go signalListen()
	
		fmt.Println("Main Start")
		Gwp.Add(1)  //控制主逻辑线程退出
		
		App = NewServerApp()
		App.Start()
		
		
		
		
		Gwp.Wait()
		fmt.Println("Main Exit")
		
		
}

func Test()  {
	println("Test")
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGKILL)
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}
