package main

import "fmt"

// interface{}是万能数据类型
// int、string、float64、struct等，全部类型都继承于它，相当于java里的Object
func myFunc(arg interface{}) {

	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 该如何区分 此时引用的底层数据类型到底是什么？

	// go 给 interface{} 提供"类型断言"的机制

	value, ok := arg.(string)

	if ok {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("value type is %T\n", value)
	} else {
		fmt.Println("arg is not string type")
	}

}

type Book struct {
	title string
}

func main() {

	myFunc(Book{"Golang"})
	myFunc(11)
	myFunc("abc")
	myFunc(12.22)

}
