package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default() //gin.Default()生成了一个实例，这个实例即 WSGI 应用程序。
//	/*使用r.Get("/", ...)声明了一个路由，告诉 Gin 什么样的URL
//	能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息*/
//	r.GET("/user/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "hello %s", name)
//	})
//	r.Run(":80") // listen and serve on 0.0.0.0:8080
//	/*
//		[GIN] 2024/03/31 - 11:52:09 | 404 |            0s |             ::1 | GET      "/"
//		[GIN] 2024/03/31 - 11:52:22 | 200 |            0s |             ::1 | GET      "/user/JJH"
//
//	*/
//}
