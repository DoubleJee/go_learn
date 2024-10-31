package main

import (
	"context"
	"fmt"
)

func print(ctx *context.Context) {
	id := (*ctx).Value("id")
	fmt.Println("id =", id)
}

func main() {
	// 创建一个带有值的 context
	ctx := context.WithValue(context.Background(), "id", "123")

	// 调用使用 context 的函数
	print(&ctx)
}
