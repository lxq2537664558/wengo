/*
创建时间: 2019/12/23
作者: zjy
功能介绍:
log 相关 数据类型定义
*/

package model

//初始化log需要的信息
type LogInitModel struct {
	 ServerName  string
	 LogsPath    string
	 LogQueueCap int //日志队列大小
}
// 日志参数
type LogModel struct {
	LogGenerateTime int64 //该条日志时间
	SceneName string
	OutStr string //具体输出的日志内容
	LogLvel int16 //日志等级
}