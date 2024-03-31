package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default() //gin.Default()生成了一个实例，这个实例即 WSGI 应用程序。
//
//	r.POST("/form", func(c *gin.Context) {
//		username := c.PostForm("username")
//		password := c.DefaultPostForm("password", "000000") // 可设置默认值
//
//		c.JSON(http.StatusOK, gin.H{
//			"username": username,
//			"password": password,
//		})
//	})
//	r.Run(":80") // listen and serve on 0.0.0.0:8080
//	/*
//		[GIN] 2024/03/31 - 11:52:09 | 404 |            0s |             ::1 | GET      "/"
//		[GIN] 2024/03/31 - 11:52:22 | 200 |            0s |             ::1 | GET      "/user/JJH"
//
//	*/
//}
