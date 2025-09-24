package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/miver02/Learn/go/webook/internal/repository"
	"github.com/miver02/Learn/go/webook/internal/web"
)


func main() {
	// 创建一个默认的 HTTP 服务器实例
	api := gin.Default()

	// 数据库层
	db := repository.NewInitDatebase().InitDB()
	rdb := repository.NewInitDatebase().InitRedis()
	repository.NewInitDatebase().InitRateLimit(api)	// 注册redis限流
	api.Use(sessions.Sessions("mysession", rdb))

	// 网络层: 跨域; 会话;
	web.NewMiddlewareBuilder().InitCors(api)
	web.NewMiddlewareBuilder().InitSess(api)

	// 注册登录效验
	// web.NewMiddlewareBuilder().LoginMiddleWareSessionBuilder(api)
	web.NewMiddlewareBuilder().LoginMiddleWareJwtBuilder(api)

	// 路由层
	api = web.RegisterRoutes(db, api)

	// 启动地址
	api.Run(":8001")
} 



