package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {

		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}()

	// 使用range迭代，可以自动判断ok
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Goroutine exit...")
}
