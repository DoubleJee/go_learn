package lib1 //一般和文件名一致
import "fmt"

// init方法，当前包的初始化
func init() {
	fmt.Println("lib1.init() ...")
}

// 首字母大写 public函数，对外开放的函数; 首字母小写private函数，当前包内调用
// 当前lib1包提供的API
func Lib1Test() {
	fmt.Println("Lib1Test() ...")
}
