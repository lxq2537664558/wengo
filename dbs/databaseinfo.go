//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  $
package dbs


type DataBaseInfo struct {
	ip string     `json:"ip"`//地址
	dbport string  `json:"dbport"`//端口号
	dbname string `json:"dbname"`//名称
	dbusername string `json:"dbusername"`//用户名
	dbpwd    string  `json:"dbpwd"`//
}

