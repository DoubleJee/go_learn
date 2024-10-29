package main

import "fmt"

// 用结构体+方法来实现oop封装----开始

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {

	// 如果类的属性首字母大写，表示该属性是对外能够访问的，否则只能类的内部访问
	Name  string
	HP    int
	level int // 私有属性
}

/*

// (this Hero) 表示与这个Hero结构体绑定一起，构成对象方法；（方法会传入this参数，类型是Hero）
// (this Hero) 表明当前方法是属于哪个结构体的
func (this Hero) GetName() string {
    // this是一个值传递

	return this.Name
}

func (this Hero) SetName(newName string) {
	// this是一个调用该方法对象的副本，值传递

	this.Name = newName
}

func (this Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("HP =", this.HP)
	fmt.Println("Level =", this.Level)
}

*/

// 首字母大写，表示其他的包能访问这个方法
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func (this *Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("HP =", this.HP)
	fmt.Println("Level =", this.level)
}

// 用结构体+方法来实现oop封装----结束

func main() {

	// 创建一个对象
	hero := Hero{Name: "张三", HP: 100, level: 1}
	hero.Show()

	hero.SetName("王五")
	hero.Show()
}
