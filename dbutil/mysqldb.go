// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.处理mysql相关逻辑
// 2.
// 3.
package dbutil

import (
	"fmt"
	"github.com/showgo/csvdata"
)

var (
	gamedb *MySqlDBStore //游戏库
	logdb *MySqlDBStore  //日志库
)



func GetMysqlDataSourceName(dbinfo *csvdata.Dbconf) string {
	if dbinfo == nil {
		fmt.Println("dbinfo is nil")
		return ""
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbinfo.Dbusername,
		dbinfo.Dbpwd,
		dbinfo.Ip,
		dbinfo.Dbport,
		dbinfo.Dbname)
}

