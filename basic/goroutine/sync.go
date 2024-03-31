package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//
//func download(url string) {
//	fmt.Println("start to download", url)
//	time.Sleep(time.Second) // 模拟耗时操作
//	wg.Done()
//}

//func main() {
//	t := time.Now()
//	for i := 0; i < 3; i++ {
//		wg.Add(1)                             //添加计数
//		go download("a.com/" + string(i+'0')) //go
//	}
//	wg.Wait()
//	fmt.Println("Done!")
//	elapsed := time.Since(t)
//	fmt.Println("app elapsed:", elapsed)
//	//原本三秒现在并发只要一秒
//	//app elapsed: 1.011689s
//}
