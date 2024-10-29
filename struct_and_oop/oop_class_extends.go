package main

import "fmt"

// 父类
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// 子类
type SuperMan struct {
	Human // 表示SuperMan类继承了Human类

	level int
}

// 重写父类的Walk方法
func (this *SuperMan) Walk() {
	fmt.Println("SuperMan.SuperMan()...")
}

// 子类的方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name =", this.name)
	fmt.Println("sex =", this.sex)
	fmt.Println("level =", this.level)
}

func main() {

	human := Human{"张三", "男"}

	human.Eat()
	human.Walk()

	fmt.Println("==================================")

	// 定义一个子类对象
	// superMan := SuperMan{Human{"超人", "男"}, 99}
	var superMan SuperMan
	// 直接.点出来父类的属性
	superMan.name = "超人"
	superMan.sex = "男"
	superMan.level = 99

	superMan.Eat()  // 父类的方法
	superMan.Walk() // 子类的方法
	superMan.Fly()  // 子类的方法

	superMan.Print()
}
