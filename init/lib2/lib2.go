package lib2 //一般和文件名一致
import "fmt"

// init方法，当前包的初始化
func init() {
	fmt.Println("lib2.init() ...")
}

// 首字母大写 public函数，对外开放的函数; 首字母小写private函数，当前包内调用
// 当前lib2包提供的API
func Lib2Test() {
	fmt.Println("Lib2Test() ...")
}
