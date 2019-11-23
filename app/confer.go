/*
创建时间: 2019/10/17
作者: zjy
功能介绍:
*/
package app

import (
	"encoding/json"
	"fmt"
	"github.com/showgo/model"
	"github.com/widuu/goini"
	"os"
	"path"
)

type Confer interface {
	InitConf() bool
	// 加载配置文件
	LoadConf()
	// 重新加载配置文件
	Reload()
}


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

	filePtr, err := os.Open("./configs/database.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	
	var infos []model.DataBaseInfo
	var content  = make([]byte,1024)
	num,erro := filePtr.Read(content)
	if erro != nil  {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	fmt.Println("解析的字符串",string(content),num)
	// 创建json解码器
	errs := json.Unmarshal(content[:num],&infos)
	if errs != nil {
		fmt.Println("Decoder failed", errs.Error())
	} else {
		fmt.Println("Decoder success")
		fmt.Println("解析结构体 =%+v",infos)
	}
}

func (sc *ServerConf)LoadConf()  {
	sc.readConfig()
}

func (sc *ServerConf)Reload()  {
	sc.LoadConf()
}