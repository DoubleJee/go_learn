package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Money int    `json:"rmb"`
}

func main() {

	user := User{"张三", 20, 200}

	// 编码的过程  结构体 --->json
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("jsonStr : %s\n", jsonStr)

	// 解码的过程  json ---> 结构体
	// {"name":"张三","age":20,"rmb":200}
	myUser := User{}
	err = json.Unmarshal(jsonStr, &myUser)
	if err != nil {
		fmt.Printf("json unMarshal error, error = %v\n", err)
		return
	}

	fmt.Println(myUser)
}
