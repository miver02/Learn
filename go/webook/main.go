package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/web"
)


func main() {
	// 数据库层
	db := InitDB()
	rdb := InitRedis()
	
	// 创建一个默认的 HTTP 服务器实例
	api := gin.Default()

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

func InitDB() *gorm.DB {
	// 数据库连接
	db, err := gorm.Open(mysql.Open("root:root@tcp(10.101.0.95:40018)/webook?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		// panic相当于整个goroutine结束
		panic("数据库连接失败: " + err.Error())
	}
	// 创建表
	err = dao.InitTable(db)
	if err != nil {
		panic("数据库表创建失败: " + err.Error())
	}
	return db
}

func InitRedis() redis.Store {
	store, err := redis.NewStore(16, "tcp", "10.101.0.95:40019", "", "redis", []byte("secret"))
	if err != nil {
		panic("Redis 连接失败: " + err.Error())
	}
	return store
}


