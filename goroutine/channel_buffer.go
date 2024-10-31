package main

import (
	"fmt"
	"time"
)

// 有缓冲区Channel
// 如果channel为空，读取数据的goroutine会阻塞
// 如果channel已满，发送数据的goroutine会阻塞

func main() {

	// 创建一个channel，有缓冲区（队列），容量为3
	c := make(chan int, 3)

	go func() {

		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			// 向c中发送数据
			c <- i
			fmt.Printf("子go程正在运行，发送的元素=%d，len(c) = %d, cap(c) = %d\n", i, len(c), cap(c))
		}

	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		// 从c中接受数据，并赋值给num
		num := <-c
		fmt.Println("num =", num)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("主go程结束")

}
