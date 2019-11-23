/*
创建时间: 2019/11/6
作者: zjy
功能介绍:
全局变量的定义 这里不做在全局包的里面 不想app包导入全局包 全局包可以导入app包
*/
package app

import (
	log "github.com/sirupsen/logrus"
)



// 路径管理相关函数
func SetAppPath(pwd string) {
	AppPath = pwd
	log.Debug("SetAppPath = ", AppPath)
}
func GetConfingsPath() string {
	return AppPath + "/configs"
}
// 配置文件名称
func GetServerIniName() string {
	return AppPath + "/configs/" + AppKindArg.ToString() + ".ini"
}

