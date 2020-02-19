/*
创建时间: 2019/12/23
作者: zjy
功能介绍:
log 相关 数据类型定义
这个文件不放在model是因为,模块内的东西尽量放在一起
*/

package xlog

type VolatileLogModel struct {
	LogQueueCap  int    // 日志队列大小
	IsOutStd     bool   // 是否在标准输出输入
	ShowLvl      uint16 // 显示日志等级
	FileTimeSpan int    // 多少小时生成一个日志文件
}

// 初始化log需要的信息
type LogInitModel struct {
	ServerName string
	LogsPath   string
	Volatile   VolatileLogModel
}

// 日志参数
type LogModel struct {
	LogGenerateTime int64 // 该条日志时间
	SceneName       string
	OutStr          string // 具体输出的日志内容
	LogLvel         uint16 // 日志等级
}
