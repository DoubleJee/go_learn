package main

import "fmt"

func main() {

	// 1. 声明slice1是一个切片，并且初始化，默认值是1 2 3， 长度是3
	// slice1 := []int{1, 2, 3}

	// 2. 声明slice1是一个切片，但是并没有给slice1分配空间
	var slice1 []int
	// 开辟3个空间，默认值是0
	// slice1 = make([]int, 3)

	// 3. 声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0
	// var slice1 []int = make([]int, 3)

	// 4. 同上，但是通过:=推导出slice1是一个切片（常用的方式）
	// slice1 := make([]int, 3)

	//slice1[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个slice是否为空
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
