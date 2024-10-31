package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// 创建一个带有超时和值的Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.WithValue(ctx, "id", "123")
	defer cancel()

	go doSomething(&ctx)

	// 等待操作完成或超时
	<-ctx.Done()
	time.Sleep(time.Second)
	fmt.Println("主函数结束...")

}

func doSomething(ctx *context.Context) {

	for {
		select {
		case <-(*ctx).Done():
			fmt.Println("操作被取消或超时")
			return
		default:
			id := (*ctx).Value("id")
			fmt.Printf("正在为id %v 执行操作...\n", id)
			time.Sleep(1 * time.Second)
		}
	}

}
