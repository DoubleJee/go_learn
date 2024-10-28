package main

import (
	// 下划线表示匿名，使用这个包的初始化方法
	_ "learn/init/lib1"

	// 包别名，有的包名很长
	mylib2 "learn/init/lib2"
	// 可以直接调用包方法，属于导入当前包里，不需要使用别名.
	// . "learn/init/lib2"
)

// 执行顺序 main包，-> 导入的第一个包，直到当前类，是个递归
// 1. init方法
// 2. 全局const
// 3. 全局var

// 最后main函数

func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()

	mylib2.Lib2Test()

	// Lib2Test()
}
