package main

import "fmt"

// 传入一个切片
func printArray(array []int) {
	// 引用传递

	// _表示匿名的变量；你不关心的变量，你不使用不访问
	// 只关心value
	for _, value := range array {
		fmt.Println("value =", value)
	}

	array[0] = 100
}

func main() {

	// 动态数组，切片 slice
	myArray := []int{1, 2, 3}

	fmt.Printf("myArray type is %T\n", myArray)

	fmt.Println("======================")

	printArray(myArray)

	fmt.Println("======================")

	for _, value := range myArray {
		fmt.Println("value =", value)
	}
}
