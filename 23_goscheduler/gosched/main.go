package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Start")
	go func() {
		fmt.Println("Start goroutine")
		for {
			stopWorkloadsomeWorkload()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}

func stopWorkload() {
	runtime.Gosched()
}
