//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  $
package gmodel

type DataBaseInfo struct {
	Ip         string      `json:ip`// 地址
	DBport     string      `json:dbport`// 端口号
	DBname     string      `json:dbname`// 名称
	DBusername string      `json:dbusername`// 用户名
	DBpwd      string      `json:dbpwd`// 密码
}

type DataBaseInfos struct {
	Dbs  []DataBaseInfo
}

func NewDataBaseInfos() *DataBaseInfos {
	return &DataBaseInfos{}
}
