/*
创建时间: 2019/11/23
作者: zjy
功能介绍:
配置文件相关数据 模型数据从jmodel来
*/

package gproxy

import ."../gmodel"

type ConfigProxy struct {
	Dbs  []DataBaseInfo
}

func NewConfigProxy() *ConfigProxy {
	return &ConfigProxy{}
}