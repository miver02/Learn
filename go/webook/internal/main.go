package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/miver02/Learn/go/webook/internal/web"
)




func main() {
	api := web.RegisterRoutes()

	// 解决跨域问题
	api.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		// AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
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

	api.Run(":8001")
} 