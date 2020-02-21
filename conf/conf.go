/*
创建时间: 2020/2/6
作者: zjy
功能介绍:
配置相关功能
*/

package conf

import (
	"github.com/showgo/xlog"
	"github.com/widuu/goini"
	"strconv"
)

var (
	conf          *goini.Config
	VolatileModel xlog.VolatileLogModel
)


func ReadIni(iniPath string)  {
	conf = goini.SetConfig(iniPath)
	conf.ReadList()
	tem, isok := strconv.Atoi(conf.GetValue("LogConf", "LogQueueCap"))
	if isok == nil {
		VolatileModel.LogQueueCap = tem
	}
	tem, isok = strconv.Atoi(conf.GetValue("LogConf", "ShowLvl"))
	if isok == nil {
		VolatileModel.ShowLvl = uint16(tem)
	}
	isOutStd, isok := strconv.ParseBool(conf.GetValue("LogConf", "IsOutStd"))
	if isok == nil {
		VolatileModel.IsOutStd = isOutStd
	}
	tem, isok = strconv.Atoi(conf.GetValue("LogConf", "FileTimeSpan"))
	if isok == nil {
		VolatileModel.FileTimeSpan = tem
	}
}