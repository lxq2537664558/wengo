// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.
// 2.
// 3.
package gmodel

//服务器相关信息
type ServerAppInfo struct {
	//多少帧每秒
	FPS  int
	//服务器类型
	Appkind AppKind
	//是否存活
	islive bool
	//服务器名称
	ServerName string
}

// app 的网络信息
type AppNetWorkInfo struct {
	//外部访问地址
	OutAddr	string
	//端口号
	OutPort	int
	//限制最大连接数
	MaxConnet int
	//最大发送多少字节
	SendMaxSize int
	//最大接受多少字节
	RecMaxSize	int
}

