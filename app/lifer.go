//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  周期接口
package app


type Lifer interface {
	// 程序启动
	Start()
	//初始化
	init() bool
	// 程序运行
	run()
	// 关闭
	Close()
}
