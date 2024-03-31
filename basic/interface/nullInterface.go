package main

import "fmt"

// 空接口，不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。
/*
func Print(a ...interface{}) (n int, err error)
func Println(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
*/

// 定义一个空接口
type Empyt_interface interface {
}

// 定义一个入参为任意类型的函数
func getInfo(arg Empyt_interface) {
	fmt.Println("getInfo 函数.....", arg)
}

// 也可以写成如下形式，更推荐
func getInfo2(arg interface{}) {
	fmt.Println("getInfo2 函数.....", arg)
}

//func main() {
//	x := "hi"
//	getInfo2(x)
//
//	map1 := make(map[string]interface{})
//	map1["数字"] = 1
//	map1["字符串"] = "字符串"
//	map1["布尔"] = false
//	fmt.Println("map1 ...........", map1)
//}
