package main

import (
	"fmt"
	"time"
)

// 子goroutine    （如果主goroutine结束，那么子goroutine也会退出结束）
func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("new Goroutine : i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {

	// 创建一个go协程，去执行一个newTask() 任务（流程）
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}
