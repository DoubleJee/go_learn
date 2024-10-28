package main

import "fmt"

// 值传递，传的是值
func changeValue(p int) {
	p = 10
}

// 指针传递，传的是指针(引用)
// 传入一个指针类型，一个指向整型的指针类型
func changeValue2(p *int) {
	// p的内存，存储的是一个指向整型的内存地址（指针）

	// *p，操作p存储的内存地址所指向的整型内存
	*p = 10
}

// 指针传递，地址传递
func swap(x, y *int) {

	temp := *x // temp = main::x
	*x = *y    // main::x = main::y
	*y = temp  // main::y = temp
}

func main() {

	var a = 1
	changeValue(a)
	fmt.Println("a =", a)
	// 传入a的内存地址
	changeValue2(&a)
	fmt.Println("a =", a)

	x, y := 10, 20
	fmt.Println("交换前：x =", x, "y =", y)
	// 传入两个变量的地址值
	swap(&x, &y)
	fmt.Println("交换后：x =", x, "y =", y)

	// 指针类型   一级指针
	var p *int
	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	// 二级指针，存一级指针地址
	var pp **int
	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)

	fmt.Println(*pp)
	fmt.Println(**pp)

}
