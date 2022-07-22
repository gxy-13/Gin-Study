package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})

	})

	// 通用方法，处理所有get post delete put
	r.Any("/hello", func(c *gin.Context) {
		//switch c.Request.Method {
		//case "GET" : ....
		//}
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})

	})

	// 用户请求了不存在的路由
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组,支持嵌套
	// 视频的首页和详情页面
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/HH", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/index/HH",
			})

		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/index/xx",
			})

		})
	}
	r.Run()
}
