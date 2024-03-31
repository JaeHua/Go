package main

import "fmt"

// 定义接口语法
//
//	type interface_name interface {
//		method_name1([args ...arg_type]) [return_type]
//
// method_name2([args ...arg_type]) [return_type]
// method_name3([args ...arg_type]) [return_type]
// ...
// method_namen([args ...arg_type]) [return_type]
// }

// 定义一个phone接口
type Phone interface {
	call()
	seenMessage()
}

// 声明结构体来实现接口
type Huawei struct {
	name  string
	price float64
}

type Xiaomi struct {
	name  string
	price float64
}

// 华为实现所有接口里的方法
func (huawei Huawei) call() {
	fmt.Printf("%s 有打电话功能.....\n", huawei.name)
}

func (huawei Huawei) seenMessage() {
	fmt.Printf("%s 有发短信功能.....\n", huawei.name)
}
func (xiaomi Xiaomi) seenMessage() {
	fmt.Printf("%s 只有发短信功能.....\n", xiaomi.name)
}

//func main() {
//	//初始化结构体，测试接口
//	mate30 := Huawei{
//		name:  "Mate 30",
//		price: 6999,
//	}
//	mate30.call()
//	mate30.seenMessage()
//	xiaomi9 := Xiaomi{
//		name:  "Xiao 9",
//		price: 4999,
//	}
//	xiaomi9.seenMessage()
//}
