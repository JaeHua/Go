package main

import "fmt"

func main() {
	//直接声明一组
	const (
		e  = 2.71828182845904523536028747135266249775724709369995957496696763
		pi = 3.14159265358979323846264338327950288419716939937510582097494459
	)
	//甚至可以省略，复制前面的
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"
	//iota 常量生成器
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday)

	//变量转换
	var f float64 = 3 + 0i // untyped complex -> float64
	f = 2                  // untyped integer -> float64
	f = 1e123              // untyped floating-point -> float64
	f = 'a'                // untyped rune -> float64
	fmt.Printf("%f", f)
}
