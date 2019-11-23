/*
创建时间: 2019/11/23
作者: zjy
功能介绍:
队列结构,封装列表集合
*/

package gcontainer

import "container/list"

type Queue struct {
	qlist  *list.List
}

// 队列构造函数
// return 返回队列
func NewQueue() (queue *Queue) {
	queue = new(Queue)
	queue.qlist = list.New()
	return queue
}

//向队列中添加数据
func (q *Queue)PushBack(v interface{})  {
	q.qlist.PushBack(v)
}

func (q *Queue)Len()  int {
	return q.qlist.Len()
}
