package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// 创建一个可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 在一个新的 goroutine 中运行一个长时间操作
	go func() {

		for {
			select {
			case <-ctx.Done():
				fmt.Println("操作被取消...")
				return
			default:
				fmt.Println("正在执行操作...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 让操作运行一段时间
	time.Sleep(2 * time.Second)

	// 取消操作
	cancel()

	// 等待 goroutine 结束
	time.Sleep(1 * time.Second)
	fmt.Println("main 协程结束...")
}
