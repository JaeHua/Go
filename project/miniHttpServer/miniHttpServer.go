package main

import (
	"fmt"
	"net/http"
)

/*
*
三行代码搭建http的web服务
http://127.0.0.1:8080
*/
func main() {
	// 设置静态文件服务,默认当请求"/"时返回当前目录下静态文件
	http.Handle("/", http.FileServer(http.Dir(".")))
	//启动HTTP服务监听端口8080,访问
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
