/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
登录服
*/

package appclient

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

type AppClient struct {
	NetWorkInfo *model.AppNetWorkModel // 服务器网络信息
	AppInfo     *model.AppInfoModel    // 服务器信息
	conns       sync.Map
	tcpclient   *network.TCPClient
}

// 程序启动
func (this *AppClient) OnStart() {
	csvdata.InitLoginCsvData()
	dispSys := dispatch.NewDispatchSys()
	dispSys.SetObserver(this)
	this.tcpclient = network.NewTCPClient(dispSys, proxy.AppConf)
	this.tcpclient.Start()
}

// 初始化
func (this *AppClient) OnInit(params interface{}) bool {
	
	return true
}

// 程序运行
func (this *AppClient) OnUpdate() bool {
	// xlog.DebugLog("","Run LoginApp")
	
	return true
}

// 关闭
func (this *AppClient) OnRelease() {
	this.tcpclient.Close()
}

func (this *AppClient) OnNetWorkConnect(conn network.Conner) error {
	xlog.DebugLog(proxy.GetSecenName(), "OnNetWorkConnect %v", conn.RemoteAddr())
	
	//连接成功发送登陆命令
	reqCreatePlyer := &login_proto.ReqLoginMsg{
		Username: "zjy0822",
		Password: "jiaopi100",
	}
	sendMsg, erro := proto.Marshal(reqCreatePlyer)
	if erro != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "OnNetWorkConnect %v", erro.Error())
		return erro
	}
	senddata,err:= conn.GetOneMsgByteArr(cmdconst.Main_LoginSv, cmdconst.Sub_C_LS_RegisterAccount, sendMsg)
	if err != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "GetOneMsgByteArr", erro)
		return erro
	}
	conn.WriteMsg(senddata,senddata)
	return nil
}

func (this *AppClient) OnNetWorkClose(conn network.Conner) error {
	xlog.DebugLog(proxy.GetSecenName(), "OnNetWorkClose %v", conn.RemoteAddr())
	return nil
}

func (this *AppClient) OnNetWorkRead(msgdata *network.MsgData) error {
	return  hanlerRead(msgdata.Conn,msgdata.MainCmd,msgdata.SubCmd,msgdata.Msgdata)
}



func hanlerRead(conn network.Conner,maincmd,subcmd uint16,msgdata []byte) error{
	xlog.ErrorLog(proxy.GetSecenName(), "AppClient hanlerRead")
	recMsg := &common_proto.RestInt32CodeMsg{}
	erro := proto.Unmarshal(msgdata,recMsg)
	if erro != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "hanlerRead %s", erro.Error())
		return erro
	}
	xlog.DebugLog(proxy.GetSecenName(),"maincmd = %d,subcmd =%d, data=%v ",maincmd,subcmd,recMsg.ResCode)
	return  nil
}
