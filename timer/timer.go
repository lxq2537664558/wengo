/*
创建时间: 2020/3/29
作者: zjy
功能介绍:

*/

package timer

import (
	"fmt"
	"time"
)

type TimerCallBack func(param interface{})

type Timer struct {
	timer *time.Timer
	timerCb  TimerCallBack
}

func NewTimer(duration time.Duration, ticb TimerCallBack) *Timer {
	timer := &Timer{
		timer:time.NewTimer(duration),
		timerCb:ticb,
	}

	go func() {
		for {
			select {
			case 	<-timer.timer.C:
				fmt.Println("Exctue time")
			}
		}
		
	}()
	return timer
}