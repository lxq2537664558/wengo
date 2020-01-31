/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
各个App的接口
*/

package app


type AppBehavior interface {
	// 程序启动
	StartApp()
	//初始化
	InitApp() bool
	// 程序运行
	RunApp()
	// 关闭
	QuitApp()
}
