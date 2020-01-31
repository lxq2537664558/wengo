/*
创建时间: 2019/12/25
作者: zjy
功能介绍:
路径相关管理
*/

package proxy

import "path"

type PathProxy struct {
	AppRootPath string //  程序(main)根路径
	CsvPath     string
	LogsPath    string
	ConfPath    string
	ConfIniPath string
	DBJsonFile  string
}

// 创建AppProxy
func NewPathProxy() *PathProxy {
	return &PathProxy{}
}

// 路径管理相关函数
func (pthpro *PathProxy) SetAppPath(pwd string) {
	pthpro.AppRootPath = pwd
	pthpro.InitProxy()
}

func (pthpro *PathProxy) InitProxy() {
	pthpro.ConfPath = path.Join(pthpro.AppRootPath, "configs")
	pthpro.CsvPath = path.Join(pthpro.LogsPath, "csv")
	pthpro.LogsPath = path.Join(pthpro.CsvPath, "logs")
	pthpro.ConfIniPath = path.Join(pthpro.ConfPath, AppPxy.AppInfo.AppKindArg.ToString()+".ini")
	pthpro.DBJsonFile = path.Join(pthpro.ConfIniPath,"database.json")
}

func (pxy *PathProxy) RealseProxy() {

}

