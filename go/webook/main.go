package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/web"

)


func main() {
	// 初始化数据库
	db := InitDB()
	
	// 创建一个默认的 HTTP 服务器实例
	api := gin.Default()

	// 为实例添加解决跨域问题功能
	api = InitCors(api)

	// 注册路由,初始化数据库
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

func InitCors(api *gin.Engine) *gin.Engine {
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