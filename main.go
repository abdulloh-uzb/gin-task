package main

import (
	"v1/config"

	"v1/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Init() {
	config.ConnectToDB()
}
func main() {
	Init()

	r := gin.Default()

	r.POST("/product", controllers.CreateProduct)

	r.Run()
}
