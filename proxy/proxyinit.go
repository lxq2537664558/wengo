/*
创建时间: 2019/12/25
作者: zjy
功能介绍:
数据代理初始化
*/

package proxy

var (
	PathPxy   *PathProxy   // 路径相关处理
	AppPxy    *AppProxy    // APP需要相关数据
)

func init() {
	createProxy()
}
// 创建代理对象
func createProxy()  {
	//创建对象在前
	AppPxy = NewAppProxy()
	PathPxy = NewPathProxy()
}
//初始化代理对象
func InitProxy()  {
	AppPxy.InitProxy()
	PathPxy.InitProxy()
}


func RealseProxy()  {
	//创建对象在前
	PathPxy = nil
	AppPxy = nil
}