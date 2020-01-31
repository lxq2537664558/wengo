/*
创建时间: 2019/11/23
作者: zjy
功能介绍:
队列结构,对标准库列表封装
*/

package queue

import (
	"container/list"
	"errors"
)

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

func (q *Queue)Front()  *list.Element{
	return q.qlist.Front()
}

func (q *Queue)PopFront() (interface{}, error)  {
	if q.Len()  == 0 {
		return  nil,errors.New("Queue Empty")
	}
	return q.qlist.Remove(q.qlist.Front()),nil
}

func (q *Queue)Clear() {
	var next *list.Element
	for elem:= q.qlist.Front(); elem != nil; elem = next {
		next = elem.Next()
		q.qlist.Remove(elem)
	}
}