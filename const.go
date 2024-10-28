package main

import "fmt"

// const 可以用来定义枚举类型
const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

const (

	// 可以在const，可以添加一个关键字 iota，每行的iota都会累加1，第一行的iota的默认值是0，下面的枚举会默认跟上面的表达式
	LOCAL  = 10 * iota // iota = 0
	REMOTE             // iota = 1
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4
	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3, i = 8, i = 12
)

func main() {

	// 常量 只读属性
	const length = 10

	fmt.Println("length =", length)

	// length = 100

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)
	fmt.Println("LOCAL =", LOCAL)
	fmt.Println("REMOTE =", REMOTE)

	fmt.Printf("a = %d\nb = %d\n", a, b)
	fmt.Printf("c = %d\nd = %d\n", c, d)
	fmt.Printf("e = %d\nf = %d\n", e, f)
	fmt.Printf("g = %d\nh = %d\n", g, h)
	fmt.Printf("i = %d\nk = %d\n", i, k)

	// iota只能出能配合const() 一起使用，iota只能在const中进行累加效果。
	// var a int = iota
}
