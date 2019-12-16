package main

import (
	"fmt"
	"reflect"
)

// 定义一个接口
type People interface {
	getAge() int     // 定义抽象方法1
	getName() string // 定义抽象方法2
}
type People2 interface {
	getAge1() int     // 定义抽象方法1
	getName1() string // 定义抽象方法2
}

type Man struct {
}

func (a *Man) getAge() int { // 实现抽象方法1
	return 18
}

func TestPeople(p interface{}) {
	fmt.Println(reflect.TypeOf(p))
	switch p.(type) { // 变量.(type) 只能在 switch 中使用
	case People:

		fmt.Println("实现了 People 接口")
		break
	case People2:
		fmt.Println("实现了 People2 接口")
		break
	default:
		fmt.Println("111")
	}
}

func main() {
	man1 := Man{}

	TestPeople(man1)
}
