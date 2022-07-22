package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取uri ，返回的都是string类型

func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSONP(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSONP(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})

	})

	r.Run(":8081")

}
