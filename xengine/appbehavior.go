/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
各个App的接口
*/

package xengine

type AppBehavior interface {
	StartApp()     // 程序启动
	InitApp() bool // 初始化
	RunApp()       // 程序运行
	QuitApp()      // 关闭
}
