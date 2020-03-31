/*
创建时间: 2020/3/29
作者: zjy
功能介绍:

*/

package main

import (
	"fmt"
	"github.com/showgo/xcontainer/queue"
	"reflect"
	"sync"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)
var queue1 *queue.Queue
var wg sync.WaitGroup

var pool sync.Pool

type Person struct {
	name string
	age int
}

func main() {
	TestPool()
}

func TestPool()  {
	pool.New = func() interface{} {
		return new(Person)
	}
	data := pool.Get()
	fmt.Println(data)
}

func TestCond(){
	queue1 = queue.NewQueue()
	_, err := queue1.PopFront()
	if err != nil {
		fmt.Println(err)
	}
	wg.Add(20)
	// 10个消费
	for i := 0; i < 10; i++ {
		go func(x int) {
			defer wg.Done()
			cond.L.Lock()         // 获取锁
			if queue1.Len() == 0 {
				cond.Wait() // 等待通知，阻塞当前 goroutine
			}
			cond.L.Unlock() // 释放锁
			val, erro := queue1.PopFront()
			if erro != nil {
				fmt.Println(erro)
				return
			}
			// do something. 这里仅打印
			fmt.Println("队列的值 ",val,"type=",reflect.TypeOf(val))
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(x int) {
			defer wg.Done()
			queue1.PushBack(x)
			cond.Signal()   // 通知其他线程
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("end")
}