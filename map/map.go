package main

import "fmt"

func main() {

	// =====> 第1种声明方式

	// 定义map1是一个map类型，key是string，value是string
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 是一个空map")
	}

	// 使用map前，需要先用make给map分配数据空间，也会自动扩容
	map1 = make(map[string]string, 10)
	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "python"

	fmt.Println(map1)

	// =====> 第二种声明方式

	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "python"
	fmt.Println(map2)

	// =====> 第三种声明方式
	map3 := map[int]string{
		1: "java",
		2: "c++",
		3: "python",
	}

	fmt.Println(map3)

}
