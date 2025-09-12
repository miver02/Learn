package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello post")
	})

	api.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "参数路由 "+name)
	})

	api.GET("/view/*.html", func(c *gin.Context) {
		page := c.Param(".html")
		c.String(http.StatusOK, "通配符路由 "+page)
	})

	api.GET("/user", func(c *gin.Context) {
		phone := c.Query("phone")
		c.String(http.StatusOK, "参数路由 "+phone)
	})

	api.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080

}
