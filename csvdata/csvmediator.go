/*
创建时间: 2020/2/11
作者: zjy
功能介绍:

*/

package csvdata

import (
	"github.com/showgo/xlog"
	"github.com/showgo/xutil"
)

var csvPath string

func SetCsvPath(csvpath string)  {
	if xutil.StringIsNil(csvpath) {
		xlog.ErrorLog("app","csvpath is nil")
	}
	csvPath  = csvpath
}

type setFunc func(csvpath string)


func ReLoadPublicCsvData()  {
	go func() {
		LoadPublicCsvData()
	}()
}
func LoadPublicCsvData()  {
	initCsvData(commoncsvset)
}

//初始化
func initCsvData(setfu []setFunc)  {
	for _,fun := range setfu {
		fun(csvPath)
	}
}
//登陆服csv相关方法
var commoncsvset = []setFunc{
	SetServerconfMapData,
	SetDbconfMapData,
}