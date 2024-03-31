package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default() //gin.Default()生成了一个实例，这个实例即 WSGI 应用程序。
//	/*使用r.Get("/", ...)声明了一个路由，告诉 Gin 什么样的URL
//	能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息*/
//	r.POST("/upload", func(c *gin.Context) {
//		file, _ := c.FormFile("file")
//		//发送的文件会下载过来
//		err := c.SaveUploadedFile(file, file.Filename)
//		if err != nil {
//			return
//		}
//		c.String(http.StatusOK, "%s uploaded!", file.Filename)
//	})
//	r.Run(":80") // listen and serve on 0.0.0.0:8080
//	/*
//		GIN] 2024/03/31 - 20:25:04 | 500 |      3.1017ms |             ::1 | POST     "/upload"
//		[GIN] 2024/03/31 - 20:25:17 | 200 |       517.5µs |             ::1 | POST     "/upload"
//
//	*/
//}
