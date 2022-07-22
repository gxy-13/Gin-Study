package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//})
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "http://cn.bing.com")
	})

	// 请求转发
	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的uri修改
		r.HandleContext(c)        // 继续后续的处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})

	})

	r.Run()

}
