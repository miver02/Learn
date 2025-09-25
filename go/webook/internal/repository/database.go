package repository

import (
	"time"

	"github.com/miver02/Learn/go/webook/pkg/ginx/middleware/ratelimit"
	redisClient "github.com/redis/go-redis/v9"

	redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/miver02/Learn/go/webook/internal/repository/dao"
	"github.com/miver02/Learn/go/webook/internal/consts"
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
func (idb *InitDatebase) InitRedis() redis.Store {
	store, err := redis.NewStore(16, "tcp", consts.RedisAddr, consts.RedisUser, consts.RedisPassword, []byte(consts.KeyPairs))
	if err != nil {
		panic("Redis 连接失败: " + err.Error())
	}
	return store
}

// 实现redis限流
func (idb *InitDatebase) InitRateLimit(api *gin.Engine) {
	rc := redisClient.NewClient(&redisClient.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPassword,
	})
	api.Use(ratelimit.NewBuilder(rc, time.Second, 100).Build())

}
