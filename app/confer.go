/*
创建时间: 2019/10/17
作者: zjy
功能介绍:
*/
package app

import "github.com/widuu/goini"

type Confer interface {
	// 加载配置文件
	LoadConf()
	// 重新加载配置文件
	Reload()
}


type ServerConf struct {
	conf *goini.Config
}

func (sc *ServerConf)load()  {
	sc.conf = goini.SetConfig(GetConfingsPath() + "\\" + GetConfigFileName())
	sc.conf.ReadList()
}

func (sc *ServerConf)LoadConf()  {
	sc.load()
}

func (sc *ServerConf)Reload()  {
	sc.load()
}