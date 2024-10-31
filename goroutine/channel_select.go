package main

import (
	"fmt"
)

// 单流程下，一个go程只能监控一个channel的状态，select可以完成监控多个channel的状态
// 有点像多路复用

func main() {

	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}

		quit <- 1
	}()

	for {

		// 监控多个channel状态
		select {
		case data := <-c:
			fmt.Println("有数据进来了, data =", data)
		case <-quit:
			fmt.Println("数据传递结束了")
			return
		}
	}

}
