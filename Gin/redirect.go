package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default()
//	//GET方法进行重定向
//	r.GET("/t", func(context *gin.Context) {
//		context.Redirect(http.StatusMovedPermanently, "/route2")
//	})
//	//通过 POST 方法进行 HTTP 重定向
//	r.POST("/test1", func(c *gin.Context) {
//		c.Redirect(http.StatusFound, "/route2")
//	})
//	//路由重定向，使用 HandleContext
//	r.GET("/route", func(c *gin.Context) {
//		c.Request.URL.Path = "/route2"
//		r.HandleContext(c)
//	})
//	r.GET("/route2", func(c *gin.Context) {
//		c.JSON(200, gin.H{"hello": "world!"})
//	})
//	r.Run()
//}
