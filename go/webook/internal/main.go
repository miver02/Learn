package main

import "github.com/miver02/Learn/go/webook/internal/web"


func main() {
	api := web.RegisterRoutes()


	api.Run(":8001")
}