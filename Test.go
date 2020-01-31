/*
创建时间: 2019/11/24
作者: zjy
功能介绍:

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mychan = make(chan int,2)

type Person struct {
	name string
	age int32
}

var wg sync.WaitGroup
var goroutinenum  uint64


func RunS(num int)  {
	
	defer  wg.Done()
	for  {
		if num % 10000 == 0 {
			fmt.Println("线程",num)
		}
		time.Sleep(time.Millisecond * 10)
	}
	
}

func memConsumed()  uint64 {
	runtime.GC() //GC，排除对象影响
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)
	return memStat.Sys
}

func main() {
	
	goroutinenum = 100000
	
	
	wg.Add(int(goroutinenum))
	before := memConsumed()
	for i:= 0 ; i < int(goroutinenum) ; i++ {
		go RunS(i)
		time.Sleep(time.Millisecond * 10)
	}
	after := memConsumed()
	
	oneUse := (after - before) / goroutinenum /1000
	fmt.Printf("%.3f KB\n", oneUse )
	wg.Wait()
	
	fmt.Println("Main End")
	
}