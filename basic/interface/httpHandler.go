package main

import (
	"fmt"
	"net/http"
)

// 一个简单的例子-使用单个处理器
type myHandler struct {
}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,GO!")
}

//func main() {
//	handler := myHandler{}
//	server := http.Server{Addr: ":8080", Handler: handler}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		panic(err)
//	}
//	//在浏览器可以看到响应
//	//curl http://localhost:8080
//	/*
//		StatusCode        : 200
//		StatusDescription : OK
//		Content           : Hello,GO!
//		RawContent        : HTTP/1.1 200 OK
//		                    Content-Length: 9
//		                    Content-Type: text/plain; charset=utf-8
//		                    Date: Sun, 31 Mar 2024 02:59:21 GMT
//
//	*/
//}
