package main

import "fmt"

// 每个变量都由 type 和 value 构成; type是类型，value是值; 把这两个称为pair
// type分为static type和concrete type
// static type静态类型，是int string....，基本类型
// concrete type具体类型，是interface所指向的具体数据类型，系统Runtime运行时能看得见知道的类型

func main() {

	var a string
	// a:  pair<static type:string, value:"abcd">
	a = "abcd"

	var allType interface{}
	// allType:  pair<type:string, value:"abcd">
	allType = a

	fmt.Println(allType)

	BookTest()

}

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}
func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func BookTest() {

	// book:  pair<type:*Book, value:"Book{}地址">
	book := &Book{}

	var r Reader
	// r:  pair<type:*Book, value:"Book{}地址">
	r = book
	r.ReadBook()

	var w Writer
	// w:  pair<type:*Book, value:"Book{}地址">
	w = r.(Writer)
	w.WriteBook()

	// 无论怎么改变，pair是不变

}
