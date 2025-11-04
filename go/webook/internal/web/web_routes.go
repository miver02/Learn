// 路由层:负责注册路由
package web

import (
	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/repository"
	"github.com/miver02/Learn/go/webook/internal/repository/cache"
	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/service"
	"github.com/miver02/Learn/go/webook/internal/service/sms/memory"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, redisClient redis.Cmdable, api *gin.Engine) *gin.Engine {
	// 初始化user
	svc, codeSvc := InitUser(db, redisClient)

	// 注册用户路由
	RegisterUserRoutes(api, svc, codeSvc)

	return api
}

func InitUser(db *gorm.DB, redisClient redis.Cmdable) (*service.UserService, *service.CodeService) {
	ud := dao.NewUserDAO(db)
	userCache := cache.NewUserCache(redisClient)
	repo := repository.NewUserRepository(ud, userCache)
	svc := service.NewUserService(repo)
	codeCache := cache.NewCodeCache(redisClient)
	codeRepo := repository.NewCodeRepository(codeCache)
	smsSvc := memory.NewService()
	codeSvc := service.NewCodeService(codeRepo, smsSvc)
	return svc, codeSvc
}

func RegisterUserRoutes(api *gin.Engine, svc *service.UserService, codeSvc *service.CodeService) {
	// 注册users路由
	u := NewUserHandle(svc, codeSvc)
	ug := api.Group("/users")
	ug.POST("/signup", func(ctx *gin.Context) { u.SignUp(ctx) })
	ug.POST("/login", func(ctx *gin.Context) { u.Login(ctx) })
	// ug.POST("/login", func(ctx *gin.Context) { u.LoginJwt(ctx) })
	// ug.POST("/login", func(ctx *gin.Context) { u.LoginSession(ctx) })
	ug.POST("/edit", func(ctx *gin.Context) { u.Edit(ctx) })
	ug.GET("/profile", func(ctx *gin.Context) { u.Profile(ctx) })
	ug.POST("/logout", func(ctx *gin.Context) { u.LogoutSession(ctx) })
	ug.POST("/login_sms/code/send", func(ctx *gin.Context) { u.SendLoginSmsCode(ctx) })
	ug.POST("/login_sms", func(ctx *gin.Context) { u.LoginSms(ctx) })
}


