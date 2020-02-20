//  创建时间: 2019/10/23
//  作者: zjy
//  功能介绍:
//  组件接口
package xengine


type Behavior interface {
	OnStart()     // 启动
	OnInit(params interface{}) bool //初始化
	OnRelease()              // 关闭On
}

type ServerBehavior interface {
	Behavior
	Updater
}
