/*
创建时间: 2020/3/3
作者: zjy
功能介绍:

*/

package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/showgo/csvdata"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"github.com/showgo/xutil"
)

// 封装数据库处理
type MySqlDBStore struct {
	db     *sql.DB
	dbConf *csvdata.Dbconf
}

func NewMySqlDBStore(dbconf *csvdata.Dbconf) *MySqlDBStore {
	dbstore := new(MySqlDBStore)
	dbstore.dbConf = dbconf
	if erro := dbstore.OpenDB(); erro != nil {
		xlog.ErrorLog(proxy.GetSecenName(), "open db error: ", erro.Error())
	}
	return dbstore
}

func (this *MySqlDBStore) OpenDB() error {
	if this.dbConf == nil {
		return errors.New("dbconf is nil")
	}
	DataSoureName := GetMysqlDataSourceName(this.dbConf)
	var Erro error
	this.db, Erro = sql.Open("mysql", DataSoureName)
	if xutil.IsError(Erro) {
		return Erro
	}
	this.db.SetMaxOpenConns(this.dbConf.Maxopenconns)
	this.db.SetMaxIdleConns(this.dbConf.Maxidleconns)
	if erro := this.db.Ping(); xutil.IsError(erro) {
		this.CloseDB()
		return erro
	}
	return Erro
}

// 关闭数据库
func (this *MySqlDBStore) CloseDB() {
	this.db.Close()
}

func (this *MySqlDBStore) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return this.db.Query(query, args ...)
}

func (this *MySqlDBStore) Excute(query string, args ...interface{}) (sql.Result, error) {
	return this.db.Exec(query, args ...)
}


func (this *MySqlDBStore) CheckTableExists(dbname string, tableName string) bool {
	if this.db == nil {
		fmt.Println("dbinfo is nil")
		return false
	}
	rows, erro := this.db.Query("SELECT t.TABLE_NAME FROM information_schema.TABLES AS t WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ", dbname, tableName)
	if xutil.IsError(erro) {
		return false
	}
	if rows.Next() {
		return true
	}
	return false
}
