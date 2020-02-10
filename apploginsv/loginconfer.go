/*
创建时间: 2020/2/3
作者: zjy
功能介绍:

*/

package apploginsv

import (
	"encoding/json"
	"github.com/showgo/model"
	"github.com/showgo/proxy"
	"github.com/showgo/xlog"
	"github.com/widuu/goini"
	"os"
)

type LoginConfer struct {
	conf *goini.Config
	Dbs  []model.DataBaseModel
}
func (lc *LoginConfer)InitConf() bool{
	return lc.LoadConf()
}
// 加载配置文件
func (lc *LoginConfer)LoadConf() bool{
	return  lc.readConfig()
}
// 重新加载配置文件
func (lc *LoginConfer)Reload() {
	lc.readConfig()
}

func (lc *LoginConfer)readConfig() bool  {
	lc.readIni()
   return lc.readDataBase()
}


func (lc *LoginConfer)readIni()  {
	lc.conf = goini.SetConfig(proxy.PathPxy.ConfIniPath)
	lc.conf.ReadList()
}

func (lc *LoginConfer)readDataBase() bool {
	dbconffilePtr, err := os.Open(proxy.PathPxy.DBJsonFile)
	if err != nil {
		xlog.DebugLog("LoginConfer","Open file failed [Err:%s]", err.Error())
		return false
	}
	defer dbconffilePtr.Close()
	// 创建json解码器
	errs := json.NewDecoder(dbconffilePtr).Decode(&lc.Dbs)
	if errs != nil {
		xlog.DebugLog("LoginConfer","Decoder failed", errs.Error())
		return false
	}
	xlog.DebugLog("LoginConfer","Decoder success 解析结构体 = %v", lc.Dbs)
	return true
}
