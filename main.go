package main

import (
	"github.com/exiaohao/http-test/pkg/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/status", controller.Status)
	router.POST("/status", controller.Status)
	router.GET("/status/:statusCode", controller.Status)
	router.POST("/status/:statusCode", controller.Status)

	router.GET("/rand_status", controller.RandResult)

	router.GET("/version", controller.Version)

	router.GET("/get", controller.GetHandler)

	router.Run(":3000")
}
