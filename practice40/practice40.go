package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("child goroutine bootstrap start")
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker .")
			case <-quit:
				fmt.Println("work well .")
				ticker.Stop()
				return
			}
		}
		//fmt.Println("child goroutine bootstrap end")
	}()
	time.Sleep(5 * time.Second)
	quit <- 1
	wg.Wait()
}
