//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"flag"
	log"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	AppPath    string         // 路径
	AppKindArg AppKind        // app类型 通过外部传递参数确定
	App        Lifer     // app每个进程只有一个
	Gwp        sync.WaitGroup // 全局的等待组 控制整个进程结束标志
)

// 这里app 的初始化工作
func init() {
}

// 获取命令行启动
func GetStart() {
	println("App init")
	var intarg int
	flag.IntVar(&intarg, "appkind", 0, "请输入app类型")
	flag.Parse()
	if intarg == 0 {
		log.Debug("请输入app类型 -appkind > 0")
	}
	AppKindArg = ItoAppKind(intarg)
	// 获取当前路径程序执行路径
	exepath, erro := os.Getwd()
	if erro != nil {
		log.Debug(erro.Error())
	}
	println(AppKindArg.ToString())
	SetAppPath(exepath)
	print(GetServerIniName())

	// sc.LoadConf()
}

