/*
创建时间: 2019/12/25
作者: zjy
功能介绍:
路径相关管理
*/

package proxy

import (
	"path"
)

type PathProxy struct {
	AppRootPath string //  程序(main)根路径
	CsvPath     string
	LogsPath    string
	ConfPath    string
	ConfIniPath string
}

// 创建AppProxy
func NewPathProxy() *PathProxy {
	return new(PathProxy)
}

// 路径管理相关函数
func (pthpro *PathProxy) SetAppPath(pwd string ) {
	pthpro.AppRootPath = pwd
}

func (pthpro *PathProxy) InitProxy() {
	pthpro.ConfPath = path.Join(pthpro.AppRootPath, "configs")
	pthpro.CsvPath = path.Join(pthpro.AppRootPath, "csv")
	pthpro.LogsPath = path.Join(pthpro.AppRootPath, "logs")
	pthpro.ConfIniPath = path.Join(pthpro.ConfPath, "serverinfo.ini")
}

func (pxy *PathProxy) RealseProxy() {

}
