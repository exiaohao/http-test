package main

import "github.com/gin-gonic/gin"
import "github.com/exiaohao/http-test/pkg/controller"

func main() {
	router := gin.Default()

	router.GET("/status", controller.Status)
	router.POST("/status", controller.Status)
	router.GET("/status/:statusCode", controller.Status)
	router.POST("/status/:statusCode", controller.Status)

	router.Run(":3000")
}
