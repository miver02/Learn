package main

import (
	
	"github.com/miver02/Learn/go/webook/internal/web"
)




func main() {
	// 注册路由,初始化数据库
	api := web.RegisterRoutes()

	


	// 启动地址
	api.Run(":8001")
} 