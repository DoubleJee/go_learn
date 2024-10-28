package main

import "fmt"

// 方法名(形参 类型; 形参 类型) 返回值类型
func foo1(a string, b int) int {

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100

	return c
}

// 多返回值，匿名
// 方法名(形参 类型; 形参 类型) (返回值类型；返回值类型)
func foo2(a string, b int) (int, int) {

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 666, 777
}

// 多返回值，有形参名的
func foo3(a string, b int) (r1 int, r2 int) {

	fmt.Println("----- foo3 ------")

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----- foo4 ------")

	// r1, r2 属于foo4的形参，初始化默认值是0
	// r1, r2 作用域空间，是foo4 整个函数体{}空间
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)

	// 给有名称的返回值变量赋值
	r1 = 888
	r2 = 999

	return

}

func main() {

	c := foo1("abc", 555)

	fmt.Println("c =", c)

	ret1, ret2 := foo2("def", 888)

	fmt.Printf("ret1 = %d\nret2 = %d\n", ret1, ret2)

	ret1, ret2 = foo3("ghi", 999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo4("ghi", 999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}
