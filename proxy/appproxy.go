/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package proxy

import (
	"github.com/showgo/model"
	"sync"
)

// App 相关数据存放
type AppProxy struct {
	AppWG        sync.WaitGroup         // app进程结束标志
	AppInfo      *model.AppInfoModel    // 服务器信息
	NetWorkInfo  *model.AppNetWorkModel // 服务器网络信息
	EndFlag        *model.AtomicInt32FlagModel
}

// 创建AppProxy
func NewAppProxy() *AppProxy {
	appPro := new(AppProxy)
	return appPro
}

func (ap  *AppProxy)InitProxy(){

}

func (ap  *AppProxy)RealseProxy(){

}



