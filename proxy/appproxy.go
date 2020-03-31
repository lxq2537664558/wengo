/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package proxy

import (
	"github.com/showgo/conf"
	"github.com/showgo/csvdata"
	"github.com/showgo/model"
	"github.com/showgo/xengine"
	"os"
	"sync"
)

var (
	PathModelPtr *model.PathModel    //最先有路径对象
	AppID        int                 //serverId
	AppConf      *csvdata.Appconf // 服务器配置
	appKind      model.AppKind       // app类型 通过外部传递参数确定
	AppFactory   xengine.AppFactory
	AppWG        sync.WaitGroup      // app进程结束标志
)


func init() {
	createProxy()
}

// 创建代理对象
func createProxy()  {
	//创建对象在前
	PathModelPtr = model.NewPathModel()
}

//初始化代理对象
func InitProxy()  {
	SetAppPath()// 获取当前路径程序执行路径
	conf.ReadIni(PathModelPtr.ConfIniPath)
}

//设置app程序路径
func SetAppPath() {
	if PathModelPtr == nil {
		return
	}
	pwd, _ := os.Getwd()
	PathModelPtr.SetRootPath(pwd)
	PathModelPtr.InitPathModel()
}


func InitKind()  {
	appKind = model.ItoAppKind(AppConf.App_kind)
}

func GetAppKind()  model.AppKind {
	return appKind
}

// App 相关数据存放
func InitAppData(appFactory  xengine.AppFactory) {
	AppFactory = appFactory
}

func RealseProxy()  {
	//创建对象在前
	PathModelPtr = nil
}

func GetSecenName() string  {
	switch appKind {
	//gameserver需要区分场景
	case model.APP_GameServer:
		return AppConf.App_name
	//这些服务器器都没有场景名称
	// case model.APP_NONE,model.APP_Client,model.APP_MsgServer,model.APP_WorldServer:
	// 	return ""
	default:
		return ""
	}
	return  ""
}


