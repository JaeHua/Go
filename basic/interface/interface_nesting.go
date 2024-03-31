package main

import "fmt"

// 定义3个接口
type A interface {
	test1()
}

type B interface {
	test2()
}

// 定义嵌套接口
type C interface {
	A
	B
	test3()
}
type Person struct {
	//如果想实现接口C，那不止要实现接口C的方法，还要实现接口A，B中的方法
}

func (p Person) test1() {
	fmt.Println("test1 方法................")
}

func (p Person) test2() {
	fmt.Println("test2 方法................")
}

func (p Person) test3() {
	fmt.Println("test3 方法................")
}

//func main() {
//	var person Person = Person{}
//	// 实现 C 接口的所有方法
//	person.test1()
//	person.test2()
//	person.test3()
//}
