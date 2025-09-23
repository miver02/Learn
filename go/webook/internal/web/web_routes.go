// 路由层:负责注册路由
package web

import (

	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/repository"
	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/service"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, api *gin.Engine) *gin.Engine {
	// 初始化user
	svc := InitUser(db)

	// 注册用户路由
	RegisterUserRoutes(api, svc)

	return api
}

func InitUser(db *gorm.DB) *service.UserService {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	return svc
}

func RegisterUserRoutes(api *gin.Engine, svc *service.UserService) {
	// 注册users路由
	u := NewUserHandle(svc)
	ug := api.Group("/users")
	ug.POST("/signup", func(ctx *gin.Context) { u.SignUp(ctx) })
	ug.POST("/login", func(ctx *gin.Context) { u.Login(ctx) })
	// ug.POST("/login", func(ctx *gin.Context) { u.LoginJwt(ctx) })
	// ug.POST("/login", func(ctx *gin.Context) { u.LoginSession(ctx) })
	ug.POST("/edit", func(ctx *gin.Context) { u.Edit(ctx) })
	ug.GET("/profile", func(ctx *gin.Context) { u.Profile(ctx) })
	ug.POST("/logout", func(ctx *gin.Context) { u.LogoutSession(ctx) })
}


