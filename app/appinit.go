//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"../gproxy"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"../gmodel"
)

var (
	App        Lifer               // app每个进程只有一个
	ConfigPxy  *gproxy.ConfigProxy // 配置文件相关数据
	AppPxy     *gproxy.AppProxy    // APP需要相关数据
)

// 这里app 的初始化工作
func init() {
	InitConfig()
}

//初始化配置文件
func InitConfig()  {
	ConfigPxy  = gproxy.NewConfigProxy()
	AppPxy = gproxy.NewAppProxy()
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
	AppPxy.AppKindArg = gmodel.ItoAppKind(intarg)
	// 获取当前路径程序执行路径
	exepath, erro := os.Getwd()
	if erro != nil {
		log.Debug(erro.Error())
	}
	println(AppPxy.AppKindArg.ToString())
	SetAppPath(exepath)
	print(GetServerIniName())
	
	// sc.LoadConf()
}
