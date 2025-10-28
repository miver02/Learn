package repository

import (
	"context"
	"time"

	"github.com/miver02/Learn/go/webook/pkg/ginx/middleware/ratelimit"
	redisClient "github.com/redis/go-redis/v9"

	redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/consts"
	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InitDatebase struct {
}

func NewInitDatebase() *InitDatebase {
	return &InitDatebase{}
}

// 连接 mysql
func (idb *InitDatebase) InitDB() *gorm.DB {
	// 数据库连接
	dsn := consts.MysqlUser + ":" + consts.MysqlPassword + "@tcp(" + consts.MysqlAddr + ")/" + consts.MysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
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

// 连接redis
// @常用于需要连接池管理的场景
// @传统 Redis 客户端，基于连接池
// @返回 redis.Store
// @github.com/gomodule/redigo/redis
func (idb *InitDatebase) InitOldRedis() redis.Store {
	store, err := redis.NewStore(16, "tcp", consts.RedisAddr, consts.RedisUser, consts.RedisPassword, []byte(consts.KeyPairs))
	if err != nil {
		panic("Redis 连接失败: " + err.Error())
	}
	return store
}

// 连接redis
// @返回 *redis.Client（该库定义的客户端结构体，直接操作 Redis）
// @支持更丰富的命令和异步操作
// @github.com/go-redis/redis/v8 或 v9
func (idb *InitDatebase) InitNewRedis() *redisClient.Client {
    client := redisClient.NewClient(&redisClient.Options{
        Addr:     consts.RedisAddr,      // Redis 地址（如 "localhost:6379"）
        Username: consts.RedisUser,      // 用户名（如果有）
        Password: consts.RedisPassword,  // 密码（如果有）
        DB:       0,                     // 默认数据库
    })
    // 测试连接
    _, err := client.Ping(context.Background()).Result()
    if err != nil {
        panic("Redis 连接失败: " + err.Error())
    }
    return client
}

// 实现redis限流
func (idb *InitDatebase) InitRateLimit(api *gin.Engine) {
	rc := redisClient.NewClient(&redisClient.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPassword,
	})
	api.Use(ratelimit.NewBuilder(rc, time.Second, 100).Build())

}
