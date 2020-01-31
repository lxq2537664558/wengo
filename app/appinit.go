//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	"../proxy"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"../model"
	"../loginserver"
)



// 这里app 的初始化工作
func init() {
	InitConfig()
}

//初始化配置文件
func InitConfig()  {

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
	AppPxy.AppInfo.AppKindArg = model.ItoAppKind(intarg)
	// 获取当前路径程序执行路径
	exepath, erro := os.Getwd()
	if erro != nil {
		log.Debug(erro.Error())
	}
	println(AppPxy.AppInfo.AppKindArg.ToString())
	SetAppPath(exepath)
	print(GetServerIniName())
	
	appBehavior = NewAppBehavior()
	
	// sc.LoadConf()
}

//app 逻辑参数根据服务器启动的参数创建对应的服务器接口
func NewAppBehavior() AppBehavior {
	switch  AppPxy.AppInfo.AppKindArg {
	case model.APP_NONE:
		return nil
	case model.APP_Client:
		return nil
	case model.APP_LogonServer: // 登录服逻辑接口
		return loginserver.NewLogionServer()
	case model.APP_GameServer:
		return nil
	case model.APP_ChatServer:
		return nil
	case model.APP_WorldServer:
		return nil
	default:
		return nil
	}
}