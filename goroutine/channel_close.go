package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {

		for i := 0; i < 3; i++ {
			c <- i
		}

		// close可以关闭一个channel
		close(c)

		// 向已关闭的channel发送数据，会报panic错误
		// c <- 1
	}()

	for {

		// ok为true，表示channel未关闭 || 有数据
		// ok为false，表示channel已关闭 && 无数据
		if data, ok := <-c; ok {
			fmt.Println("data =", data)
		} else {
			break
		}
	}

	fmt.Println("Main finished..")

}
