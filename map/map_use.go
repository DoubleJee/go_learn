package main

import "fmt"

func printMap(myMap map[string]string) {
	// myMap是一个引用传递

	for key, value := range myMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}

}

func main() {

	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	// 修改
	cityMap["USA"] = "DC"

	// 删除
	delete(cityMap, "Japan")

	fmt.Println("========================")

	printMap(cityMap)
}
