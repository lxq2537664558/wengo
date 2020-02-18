/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package proxy

import (
	"github.com/showgo/csvdata"
	"github.com/showgo/model"
	"github.com/showgo/xengine"
	"sync"
)

var (
	AppFactory    xengine.AppFactory
	AppWG         sync.WaitGroup      // app进程结束标志
	SvConf        *csvdata.Serverconf // 服务器配置
	AppKindArg    model.AppKind       // app类型 通过外部传递参数确定
	SververID     int    //serverId
)

func InitKind()  {
	AppKindArg = model.ItoAppKind(SvConf.ServerType)
}

// App 相关数据存放
func InitAppData(appFactory  xengine.AppFactory) {
	AppFactory = appFactory
}


