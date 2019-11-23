/*
创建时间: 2019/11/23
作者: zjy
功能介绍:

*/

package app

import (
	"encoding/json"
	"fmt"
	"github.com/widuu/goini"
	"os"
)

type ServerConf struct {
	conf *goini.Config
}

// 创建server对象
func NewserverConf() Confer {
	return &ServerConf{}
}

func (sc *ServerConf)InitConf() bool{
	// sc.conf = goini.SetConfig(GetConfingsPath() + "\\" + GetConfigFileName())
	// sc.conf.ReadList()
	sc.LoadConf()
	return  true
}



func (sc *ServerConf)readConfig()  {
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

func (sc *ServerConf)LoadConf()  {
	sc.readConfig()
}

func (sc *ServerConf)Reload()  {
	sc.LoadConf()
}
