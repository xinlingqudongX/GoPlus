package util

import (
	"time"
)

func init(){

}

//Delay 延迟执行函数
//handler 是执行函数
//delay	是延迟的秒数
func Delay(handler func(), delay int) {
	trigger := time.After(time.Second * time.Duration(delay))
	<-trigger
	if handler != nil {
		handler()
	}
}

//Interval 循环执行
//handler 执行的函数，返回为假时中断
//interval	间隔的秒数时间
func Interval(handler func() bool, interval int) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	for range ticker.C {
		if handler != nil && !handler() {
			break
		}
	}
}

