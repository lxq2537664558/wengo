//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  组件接口
package xengine

type XBehavior interface {
	OnStart() // 启动
	OnInit() bool//初始化
	OnRun()	   // 运行
	OnClose() // 关闭On
}
