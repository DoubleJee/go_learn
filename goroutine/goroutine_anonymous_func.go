package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 用go创建承载了一个形参为空、无返回值的匿名函数，并执行
	go func() {
		defer fmt.Println("A.Defer")

		// 定义匿名函数，并调用
		func() {
			defer fmt.Println("B.Defer")

			// 退出当前goroutine   （方法会中断相继出栈，执行defer）
			runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// 有参有返回值—匿名函数，调用
	// 因为是异步的，无法拿到返回值，go程只能通过Channel来通信
	go func(a, b int) bool {
		fmt.Println("a =", a, ", b =", b)
		return true
	}(1, 2)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
