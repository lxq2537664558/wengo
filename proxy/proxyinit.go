/*
创建时间: 2019/12/25
作者: zjy
功能介绍:
数据代理初始化
*/

package proxy

var (
	PathPxy   *PathProxy   // 路径相关处理
	ConfigPxy *ConfigProxy // 配置文件相关数据
	AppPxy    *AppProxy    // APP需要相关数据
)

func init() {
	createProxy()
	InitProxy()
}
// 创建代理对象
func createProxy()  {
	//创建对象在前
	PathPxy = NewPathProxy()
	ConfigPxy = NewConfigProxy()
	AppPxy = NewAppProxy()
}
//初始化代理对象
func InitProxy()  {
	PathPxy.InitProxy()
	ConfigPxy.InitConf()
}


func RealseProxy()  {
	//创建对象在前
	PathPxy = NewPathProxy()
	ConfigPxy = NewConfigProxy()
	AppPxy = NewAppProxy()
}