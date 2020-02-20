/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
登录服
*/

package apploginsv

import (
	"github.com/showgo/csvdata"
	"github.com/showgo/model"
	"github.com/showgo/network"
)

type LogionServer struct {
	NetWorkInfo  *model.AppNetWorkModel // 服务器网络信息
	AppInfo      *model.AppInfoModel    // 服务器信息
}


// 程序启动
func (ls *LogionServer)OnStart() {
	csvdata.InitLoginCsvData()
	tcpserver := new(network.TCPServer)
	tcpserver.Start()
}

//初始化
func (ls *LogionServer)OnInit(params interface{}) bool{

	return true
}
// 程序运行
func (ls *LogionServer)OnUpdate(){
   // xlog.DebugLog("","Run LoginApp")
}
// 关闭
func (ls *LogionServer)OnRelease(){

}