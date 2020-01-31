/*
创建时间: 2019/12/22
作者: zjy
功能介绍:

*/

package xengine


type Component interface {
	GetEntity() (*Entity,error)
}

type UpdateComponent interface {
	Update() //更新
}
