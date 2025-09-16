package web

import (
	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/repository"
	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegisterRoutes() *gin.Engine {
	api := gin.Default()
	// 注册用户路由
	RegisterUserRoutes(api)

	return api
}

func RegisterUserRoutes(api *gin.Engine) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := NewUserHandle(svc)

	ug := api.Group("/users")
	ug.POST("/signup", func(ctx *gin.Context) { u.SignUp(ctx) })
	ug.POST("/login", func(ctx *gin.Context) { u.Login(ctx) })
	ug.POST("/edit", func(ctx *gin.Context) { u.Edit(ctx) })
	ug.GET("/profile", func(ctx *gin.Context) { u.Profile(ctx) })
	// ug.POST("/logout", func(ctx *gin.Context) {})
}
