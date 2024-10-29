package main

import (
	"fmt"
)

func printArray(array [5]int) {
	// 值拷贝

	for index, value := range array {
		fmt.Println("index =", index, "value =", value)
	}

	array[0] = 111
}

func main() {

	// 固定长度
	var array [10]int

	array2 := [5]int{1, 2, 3}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// range关键字，返回数组/集合每个元素，有index和value
	for index, value := range array2 {
		fmt.Println("index =", index, "value =", value)
	}

	// 数组类型是携带长度的，[10]int 是一个类型，[5]int是一个类型
	fmt.Printf("array types = %T\n", array)
	fmt.Printf("array2 types = %T\n", array2)

	printArray(array2)
}
