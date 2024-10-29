package main

import "fmt"

// 声明一种新的数据类型 myint，是int的一个别名
type myint int

// 定义一个结构体;  多种数据类型组合到一起，变成一种复杂的数据类型，类似类Class
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	// 传递一个book的副本;  值传递

	book.author = "李四"
}

func changeBook2(book *Book) {
	// 指针传递

	// 关于*，go编译器对结构体做了优化
	book.author = "王五"
}

func main() {

	/*
		var a myint = 10
		fmt.Printf("a = %d, type of a = %T\n", a, a)
	*/

	// 创建一个结构体
	var book1 Book
	book1.title = "Golang语言开发"
	book1.author = "张三"

	fmt.Println(book1)

	changeBook(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)

	// 创建方式2
	var book2 = Book{title: "赚钱门道", author: "李四"}
	fmt.Println(book2)
}
