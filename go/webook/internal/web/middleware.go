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
	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareBuilder struct {
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{}
}

func (mb *MiddlewareBuilder) InitCors(api *gin.Engine) {
	// 解决跨域问题
	api.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		// AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"}, // 不加这个,前端拿不到token
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			if strings.HasPrefix(origin, "live.webook.com") {
				return true
			}
			return strings.Contains(origin, "xxx.com'")
		},
		MaxAge: 12 * time.Hour,
	}))
}

func (mb *MiddlewareBuilder) InitSess(api *gin.Engine) {
	// 配置会话中间件
	store := cookie.NewStore([]byte("secret"))
	api.Use(sessions.Sessions("mysession", store))
}

func (mb *MiddlewareBuilder) LoginMiddleWareSessionBuilder(api *gin.Engine) {
	// 注册登录效验
	api.Use(func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		sess := sessions.Default(ctx)
		idVal := sess.Get("UserId")
		if idVal == nil {
			// 没有登陆
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("未登录会话"))
			return
		}

		// 1. 无论是否首次访问，都显式设置 MaxAge（核心修正）
		sess.Options(sessions.Options{
			MaxAge:   10 * 60,     // 10分钟，按你的需求设置
			Path:     "/",         // 确保路径匹配，避免多个 Cookie 冲突
			Domain:   "127.0.0.1", // 明确域名，与请求一致
			HttpOnly: true,
			Secure:   false, // 本地开发关闭 Secure
		})

		// 2. 处理 update_time 逻辑（保持你的业务逻辑）
		now := time.Now().UnixMilli()
		updateTime := sess.Get("update_time")

		if updateTime == nil {
			sess.Set("update_time", now)
		} else {
			updateTimeval, _ := updateTime.(int64)
			if now-updateTimeval > 60*1000*9 { // 1分钟更新一次
				sess.Set("update_time", now)
			}
		}

		// 3. 关键：无论是否修改，都保存会话（确保 Options 生效）
		if err := sess.Save(); err != nil {
			// 处理保存错误
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("会话保存失败"))
			return
		}

		// 打印会话 ID 和 userId，查看是否为空
		fmt.Printf("会话校验:会话ID: %s, userid: %v\n", sess.ID(), idVal)
	})
}

func (mb *MiddlewareBuilder) LoginMiddleWareJwtBuilder(api *gin.Engine) {
	api.Use(func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := segs[1]
		// 提取token中的参数,一定要传入指针
		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// jwt过期了,token.Valid会变为False
		if token == nil || !token.Valid || claims.Uid == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.UserAgent != ctx.Request.UserAgent() {
			// 严重的安全问题,需要记录日志
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 刷新jwt
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*30 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err := token.SignedString([]byte("secret"))
			if err != nil {
				ctx.String(http.StatusOK, err.Error())
			}
			ctx.Header("x-jwt-token", "Bearer "+tokenStr)
		}

		ctx.Set("claims", claims)
		fmt.Printf("Jwt校验:uid: %d\n", claims.Uid)
	})
}
