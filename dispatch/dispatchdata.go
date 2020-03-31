/*
创建时间: 2020/2/29
作者: zjy
功能介绍:

*/

package dispatch

const (
	Timer_Event         = iota // 定时器事件
	NetWorkAccept_Event        // 网络连接事件
	NetWorkRead_Event          // 网络读取事件
	NetWorkClose_Event         // 网络
	DisPatch_max
)

type DisPatchData struct {
	dipatchType int
	val         interface{}
}

// 事件数据
func DisPatchDataPoolNewFun() interface{} {
	return new(DisPatchData)
}
// 事件数据
func NewDisPatchData(dtype int, val interface{}) *DisPatchData {
	data := new(DisPatchData)
	data.SetDisPatchData(dtype,val)
	return data
}


// 事件数据
func (this *DisPatchData)SetDisPatchData(dtype int, val interface{}){
	this.dipatchType = dtype
	this.val = val
}
