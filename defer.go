package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func deferFunc() int {
	fmt.Println("defer func called....")
	return 0
}

func retrunFunc() int {
	fmt.Println("return func called....")
	return 1
}

// return先执行、defer后执行，返回值还是return的
func returnAndDefer() int {

	defer deferFunc()

	return retrunFunc()
}

func main() {

	// defer关键字 方法执行完毕后执行的表达式（在return之后）
	// defer fmt.Println("main::end1")

	// fmt.Println("main::hello1")
	// fmt.Println("main::hello2")

	// defer按照LIFO栈结构执行，先定义的后执行
	// defer func1()
	// defer func2()
	// defer func3()

	v := returnAndDefer()
	fmt.Println("v =", v)

}
