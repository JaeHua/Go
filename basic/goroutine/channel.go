package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道
func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	ch <- url               // 将 url 发送给信道
}

func main() {
	t := time.Now()
	for i := 0; i < 3; i++ { //添加计数
		go download("a.com/" + string(i+'0')) //go download()：启动新的协程并发执行 download 函数
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
	//原本三秒现在并发只要一秒
	//app elapsed: 1.011689s
}
