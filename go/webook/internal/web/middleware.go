package web

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type MiddlewareBuilder struct {
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{}
}

func (mb *MiddlewareBuilder) InitCors(api *gin.Engine) *gin.Engine {
	// 解决跨域问题
	api.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		// AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "xxx.com'")
		},
		MaxAge: 12 * time.Hour,
	}))

	return api
}

func (mb *MiddlewareBuilder) InitSess(api *gin.Engine) *gin.Engine {
	// 配置会话中间件
	store := cookie.NewStore([]byte("secret"))
	api.Use(sessions.Sessions("mysession", store))
	return api
}

func (mb *MiddlewareBuilder) LoginMiddleWareMiddlewareBuilder(api *gin.Engine) *gin.Engine {
	// 注册登录效验
	api.Use(func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		sess := sessions.Default(ctx)
		id := sess.Get("UserId");
		// 打印会话 ID 和 userId，查看是否为空
		fmt.Printf("会话ID: %s\n", sess.ID())
		fmt.Printf("userId: %v\n", id)
		if id == nil {
			// 没有登陆
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("未登录会话"))
			return
		}
	})
	return api
}


