//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  $
package model

type DataBaseInfo struct {
	Ip         string          // 地址
	DBport     string    // 端口号
	DBname     string      // 名称
	DBusername string // 用户名
	DBpwd      string      // 密码
}

type DataBaseInfos struct {
	Test string
	Dbs  []DataBaseInfo
}

func NewDataBaseInfos() *DataBaseInfos {
	return &DataBaseInfos{}
}
