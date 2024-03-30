package main

//import "fmt"
//
//func main() {
//	//三种方法声明
//	//var a [3]int;
//	//var a [3]int = [3]int{1, 2, 3}
//	//a := [...]int{0: 1, 1: 2, 2: 3}
//	a := [...]int{1, 2, 3}
//	//len
//	fmt.Println(a[len(a)-1])
//	//必须{放在上面
//	for i, num := range a {
//		fmt.Printf("%d : %d\n", i, num)
//	}
//
//	//类似的有
//	type Currency int
//	const (
//		USD Currency = iota
//		EUR
//		RMB
//		GBP
//	)
//	symbol := [...]string{USD: "$", EUR: "€", RMB: "￡", GBP: "￥"}
//	fmt.Println(RMB, symbol[RMB])
//}
