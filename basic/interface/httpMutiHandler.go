package main

import (
	"fmt"
	"log"
	"net/http"
)

// 多处理器
type HelloHandler struct{}

// 实现接口的方法
func (HelloHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello!")
}

type WorldHandler struct{}

func (WorldHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "World!")
}
func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/hello", hello)
	mux.Handle("/world", world)

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
