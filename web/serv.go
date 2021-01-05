package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//func hello(w http.ResponseWriter, r *http.Request) {
//if r.Method == "GET" {
//io.WriteString(w, "<html><head></head><body><h1>Welcome!</h1><img src=\"/static/IMG_0547.JPG\"></body></html>")
//return
//}
//}

func homeFunc(c *gin.Context) {
	c.String(http.StatusOK, "<html><head></head><body><h1>Welcome!</h1><img src=\"/static/IMG_0547.JPG\"></body></html>", nil)
}

func StartUp() {
	gin.SetMode(gin.ReleaseMode)
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.Static("/static", "./static")
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/home", homeFunc)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8080")
}
