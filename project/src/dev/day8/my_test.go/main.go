package main

import (
	"fmt"
)

func main() {

	// 在main中向chan写入大于容量的数据，死锁    解决：开一个协程来做
	intChan := make(chan int, 1000)

	for i := 0; i < 1001; i++ {
		intChan <- i
	}
    // 解决
	intChan := make(chan int, 1000)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
	}()







	// 在main中读chan，如果读一个空的chan，死锁     解决：开一个协程来做
	exitChan := make(chan bool, 8)
	for i := 0; i < 7; i++ {
		exitChan <- true
	}
	for i := 0; i < 8; i++ {
		<-exitChan
	}
	// 解决
	exitChan := make(chan bool, 8)
	for i := 0; i < 7; i++ {
		exitChan <- true
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
	}()

}
