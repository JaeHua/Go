package main

//import "fmt"
//
//// 值传递与指针传递
//func main() {
//	a := Number{1}
//	b := Number{1}
//	a.Add(1)
//	fmt.Printf("a: %d\n", a.sum) //a --> 1
//	b.Addptr(1)
//	fmt.Printf("b: %d\n", b.sum) //b -->1
//
//	c := a
//	c.Add(1)
//	fmt.Printf("a: %d\n", a.sum) //a --> 1
//	fmt.Printf("c: %d\n", c.sum) //c -->	1
//
//	c.Addptr(1)
//	fmt.Printf("a: %d\n", a.sum) //a --> 1
//	fmt.Printf("c: %d\n", c.sum) //c --> 2
//
//	e := a.Add(1)
//	fmt.Printf("e: %d\n", e.sum) //e-->2
//
//	d := b
//	d.Addptr(1)
//	fmt.Printf("b: %d\n", b.sum) //b --> 2
//	fmt.Printf("d: %d\n", d.sum) //d --> 3
//}
//
//type Number struct {
//	sum int
//}
//
//func (n Number) Add(i int) Number {
//	n.sum += i
//	return n
//}
//func (np *Number) Addptr(i int) {
//	np.sum += i
//}
