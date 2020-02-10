//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  $
package model

type DataBaseModel struct {
	Ip         string      `json:ip`// 地址
	DBport     string      `json:dbport`// 端口号
	DBname     string      `json:dbname`// 名称
	DBusername string      `json:dbusername`// 用户名
	DBpwd      string      `json:dbpwd`// 密码
}



