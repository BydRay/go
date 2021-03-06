package main

import (
	"time"
	"runtime"
)


func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 10240; i++ {
		go  func() {
			for {
				select {
					case <-time.After(time.Nanosecond)
				}
			}
		}()
	}

	time.Sleep(time.Second * 10000)
}