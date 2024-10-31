package main

import (
	"fmt"
	"time"
)

// 接收端，如果channel数据为空，读取数据的goroutine会阻塞
// 发送端，如果发送的数据没人接收，发送数据的goroutine会阻塞
// 是一次耦合操作，发送端等到接收端接受， 接收端等到发送端发送过来; 像一次接力棒的过程，互相等待

func main() {

	// 创建一个channel，传输int类型的channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine 正在运行...")

		// 将666 发送给channel
		c <- 666
	}()

	time.Sleep(2 * time.Second)
	// 从c中接受数据，并赋值给num
	num := <-c

	fmt.Println("num =", num)
	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine结束...")
}
