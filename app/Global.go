// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.

package app

import(
	"sync"
	"github.com/name5566/leaf/log"
)


var (
	AppPath string            // 路径
	AppKindArg AppKind	           // app类型 通过外部传递参数确定
	App  *ServerApp           // app每个进程只有一个
	Gwp  sync.WaitGroup       // 全局的
	)
// 路径管理相关函数
func SetAppPath(pwd string) {
	AppPath = pwd
	log.Debug("SetAppPath = ", AppPath)
}
func GetConfingsPath() string {
	return AppPath + "/configs"
}

//配置文件名称
func GetConfigFileName() string {
	return "ServerInfo.ini"
}
