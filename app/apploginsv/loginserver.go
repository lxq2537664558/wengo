/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
登录服
*/

package apploginsv

import (
	"github.com/golang/protobuf/proto"
	"github.com/showgo/cmdconst"
	"github.com/showgo/csvdata"
	"github.com/showgo/dispatch"
	"github.com/showgo/model"
	"github.com/showgo/network"
	"github.com/showgo/protobuf/pb/common_proto"
	"github.com/showgo/protobuf/pb/login_proto"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"sync"
)



type LogionServer struct {
	NetWorkInfo  *model.AppNetWorkModel // 服务器网络信息
	AppInfo      *model.AppInfoModel    // 服务器信息
	tcpserver    *network.TCPServer
	conns        sync.Map
	dispSys *dispatch.DispatchSys
}


// 程序启动
func (this *LogionServer)OnStart() {
	csvdata.InitLoginCsvData()
	this.dispSys = dispatch.NewDispatchSys()
	this.dispSys.SetObserver(this)
	this.tcpserver = network.NewTcpServer(this.dispSys,proxy.AppConf)
	this.tcpserver.Start()
}

//初始化
func (this *LogionServer)OnInit(params interface{}) bool{

	return true
}
// 程序运行
func (this *LogionServer)OnUpdate() bool{
	
	return true
}
// 关闭
func (this *LogionServer)OnRelease(){
	this.tcpserver.Close()
	this.dispSys.Release()
}

func (this *LogionServer)OnNetWorkConnect(conn network.Conner) error{
	xlog.DebugLog(proxy.GetSecenName(),"OnNetWorkConnect %v",conn.RemoteAddr())
	return  nil
}


func (this *LogionServer)OnNetWorkClose(conn network.Conner) error{
	xlog.DebugLog(proxy.GetSecenName(),"OnNetWorkClose %v",conn.RemoteAddr())
	return  nil
}

func (this *LogionServer)OnNetWorkRead(msgdata *network.MsgData) error{
	xlog.ErrorLog(proxy.GetSecenName(), "LogionServer OnNetWorkRead",)
	return  hanlerRead(msgdata.Conn,msgdata.MainCmd,msgdata.SubCmd,msgdata.Msgdata)
}


func hanlerRead(conn network.Conner,maincmd,subcmd uint16,msgdata []byte) error{
	reqCreatePlyer := &login_proto.ReqLoginMsg{}
	erro := proto.Unmarshal(msgdata,reqCreatePlyer)
	if erro != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "OnNetWorkConnect %v", erro.Error())
		return erro
	}
	xlog.DebugLog(proxy.GetSecenName(),"maincmd = %d,subcmd =%d, data=%s,%s ",maincmd,subcmd,reqCreatePlyer.Username,reqCreatePlyer.Password)
	
	restcode := &common_proto.RestInt32CodeMsg{
		ResCode:100,
	}
	//收到消息回复
	sendMsg, erro := proto.Marshal(restcode)
	if erro != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "hanlerRead %s", erro.Error())
		return erro
	}
	conn.WriteOneMsg(cmdconst.Main_LoginSv, cmdconst.Sub_C_LS_RegisterAccount, sendMsg)
	
	return  nil
}


