package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

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

	// 等待操作完成或超时
	<-ctx.Done()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("主函数结束")
}
