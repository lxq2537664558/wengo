/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/widuu/goini"
	"os"
	."../model"
	"path"
)

type ConfigProxy struct {
	conf *goini.Config
	Dbs  []DataBaseModel
}

// 创建server对象
func NewConfigProxy() Confer {
	return &ConfigProxy{}
}
func (pthpro *ConfigProxy) InitProxy() {

}
func (pxy *ConfigProxy) RealseProxy() {

}

func (sc *ConfigProxy)InitConf() bool{
	sc.readConfig()
	return  true
}

func (sc *ConfigProxy)readConfig()  {
	dbconffilePtr, err := os.Open(GetDBJsonFile())
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer dbconffilePtr.Close()
	// 创建json解码器
	errs := json.NewDecoder(dbconffilePtr).Decode(&ConfigPxy.Dbs)
	if errs != nil {
		fmt.Println("Decoder failed", errs.Error())
	} else {
		fmt.Println("Decoder success")
		fmt.Println("解析结构体 =%+v",ConfigPxy.Dbs)
	}
}

func (sc *ConfigProxy)readIni()  {
	sc.conf = goini.SetConfig(GetConfingsPath() + "\\" + GetConfigFileName())
	sc.conf.ReadList()
}
func (sc *ConfigProxy)readDataBase()  {

}

func (sc *ConfigProxy)LoadConf()  {

}

func (sc *ConfigProxy)Reload()  {
	sc.LoadConf()
}
