package main

import "fmt"

// 声明全局变量 方法一、二、三是可以的
var gA int = 100
var gB = 200

// 方法四来声明全局变量
// := 只能够在函数体内来声明，局部变量
// gC := 300

func main() {

	// 方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三：自动匹配类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四：（常用方法）省去var关键字，直接自动匹配
	d := "love"
	fmt.Printf("type of d = %T\n", d)

	g := 3.14
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA =", gA, "gB =", gB)
	// fmt.Println("gC =", gC)

	// 声明多个变量
	var xx, yy int = 11, 22
	fmt.Println("xx =", xx, "yy =", yy)

	var kk, ll = 100, "abc"
	fmt.Println("kk =", kk, "ll =", ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, "jj =", jj)

}
