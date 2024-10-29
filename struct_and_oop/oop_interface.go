package main

import (
	"fmt"
)

// 定义接口;  本质是一个指针
type AnimalIF interface {

	// 休息
	Sleep()

	// 获取动物的颜色
	GetColor() string

	// 获取动物的种类
	GetType() string
}

// 子类;  实现接口，无需在类里写接口类型定义，只要你实现了接口全部方法，你就是一个AnimalIF接口实现类
type Cat struct {
	color string
}

// Car实现了AnimalIF接口的全部方法
func (this *Cat) Sleep() {
	fmt.Println("猫.sleep()...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "猫"
}

// 子类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("狗.sleep()...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "狗"
}

// 传入接口，接口本质是一个指针，指向实现子类地址
func showAnimal(animal AnimalIF) {

	// 多态
	animal.Sleep()

	fmt.Println("color =", animal.GetColor())
	fmt.Println("type =", animal.GetType())
}

func main() {

	/*
		var animal AnimalIF

		animal = &Cat{"黑色"}
		animal.Sleep() // 调用的就是Cat的Sleep方法

		animal = &Dog{"黄色"}
		animal.Sleep() // 调用的就是Dog的Sleep方法
	*/

	cat := Cat{"黑色"}
	dog := Dog{"黄色"}

	showAnimal(&cat)
	showAnimal(&dog)

}
