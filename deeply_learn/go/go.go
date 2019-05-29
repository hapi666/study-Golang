package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// 第一种
	var count uint32
	tigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1) //原子操作
				break
			}
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			tigger(i, fn)
		}(i)
	}
	tigger(10, func() {})
	// 第二种
	// ch := make(chan struct{}, 10)
	for index := 0; index < 10; index++ {
		go func(int) {
			// fmt.Println(index)
			// ch <- struct{}{}
		}(index)
		// <-ch
		// ch <- struct{}{}
	}
	// <-ch
	time.Sleep(time.Nanosecond)
}
