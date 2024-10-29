package main

import "fmt"

func main() {

	// len = 3, cap = 3
	s := []int{4, 5, 6}

	// 截取切片，左闭右开 [0, 2)，浅拷贝
	s1 := s[0:2]
	fmt.Println("s1 =", s1)

	// s1和s指向的都是一块数组内存地址，但是s1通过len来控制了可操作的范围
	s1[0] = 100
	fmt.Println("s =", s)
	fmt.Println("s1 =", s1)

	// copy 可以将底层数组的slice一起进行拷贝
	var s2 []int = make([]int, 3) // s2 = [0, 0, 0]
	// 将s中的值，依次拷贝到s2中
	copy(s2, s)

	fmt.Println("s2 =", s2)

	// 截取全部
	fmt.Println(s[:])
	// 截取 从0开始到2
	fmt.Println(s[:2])
	// 截取 从1开始到末尾
	fmt.Println(s[1:])
}
