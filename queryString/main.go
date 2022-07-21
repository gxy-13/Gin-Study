package main

// queryString
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		// 获取客户端请求的参数

		// 通过query 获取请求中携带的queryString参数
		name := c.Query("query")
		age := c.Query("age")

		// 没有取到query的参数，就用默认值somebody
		//name := c.DefaultQuery("query", "somebody")

		//// 取到返回(值,true)    没有取到("",false)
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	// 没取到
		//	name = "somebody"
		//}
		c.JSONP(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8090")
}
