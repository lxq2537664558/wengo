/*
创建时间: 2019/10/17
作者: zjy
功能介绍:
*/
package xengine

type Confer interface {
	InitConf() bool
	LoadConf() bool// 加载配置文件
	Reload()  // 重新加载配置文件
}
