package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	api := gin.Default()
	// 注册用户路由
	RegisterUserRoutes(api)

	return api
}

func RegisterUserRoutes(api *gin.Engine) {
	u := NewUserHandle()

	ug := api.Group("/users")
	ug.POST("/signup", func(ctx *gin.Context) { u.SignUp(ctx) })
	ug.POST("/login", func(ctx *gin.Context) { u.Login(ctx) })
	ug.POST("/edit", func(ctx *gin.Context) { u.Edit(ctx) })
	ug.GET("/profile", func(ctx *gin.Context) { u.Profile(ctx) })
	// ug.POST("/logout", func(ctx *gin.Context) {})
}
