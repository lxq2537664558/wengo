/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
登录服
*/

package loginserver


type LogionServer struct {

}
// 创建一个服务器
func NewLogionServer()  *LogionServer {
	ls := new(LogionServer)
	return ls
}

// 程序启动
func (ls *LogionServer)StartApp() {

}
//初始化
func (ls *LogionServer)InitApp() bool{

	return true
}
// 程序运行
func (ls *LogionServer)RunApp(){

}
// 关闭
func (ls *LogionServer)QuitApp(){

}