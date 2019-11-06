//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  app 初始化工作
package app

import (
	log "github.com/sirupsen/logrus"
	"flag"
	"os"
)

// 这里app 的初始化工作
func init() {
	println("App init" )
	var intarg int
	flag.IntVar(&intarg,"AppKind", 0,"请输入app类型")
	flag.Parse()
	if intarg == 0 {
		log.Debug("请输入app类型 -AppKind > 0" )
	}
	AppKindArg =  ItoAppKind(intarg)
	//获取当前路径程序执行路径
	exepath,erro := os.Getwd()
	if  erro != nil{
		log.Debug(erro.Error())
	}
	println(AppKindArg.String())
	SetAppPath(exepath)
}
