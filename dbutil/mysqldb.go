// 创建时间: 2019/10/17
// 作者: zjy
// 功能介绍:
// 1.处理mysql相关逻辑
// 2.
// 3.
package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/showgo/csvdata"
	"github.com/showgo/xutil"
)

func OpenDB(dbinfo *csvdata.Dbconf) *sql.DB {
	if dbinfo == nil {
		return nil
	}
	DataSoureName := GetMysqlDataSourceName(dbinfo)
	// gorm.Open()
	db,Erro := sql.Open("mysql", DataSoureName)
	if xutil.IsError(Erro) {
		return nil
	}
	db.SetMaxOpenConns(dbinfo.Maxopenconns)
	db.SetMaxIdleConns(dbinfo.Maxidleconns)
	if erro := db.Ping() ; xutil.IsError(erro) {
		db.Close()
		return nil
	}
	return db
}

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

func CheckTableExists(db *sql.DB,dbname string,tableName string) bool{
	if db == nil {
		fmt.Println("dbinfo is nil")
		return false
	}
	rows,erro := db.Query("SELECT t.TABLE_NAME FROM information_schema.TABLES AS t WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ",dbname,tableName)
	if xutil.IsError(erro) {
		return false
	}
	if rows.Next() {
		return true
	}
	return false
}