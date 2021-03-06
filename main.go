package main

import (
	"os"

	"github.com/exiaohao/http-test/pkg/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	runPort := os.Getenv("HTTP_PORT")
	router := gin.Default()

	router.GET("/", controller.Welcome)
	router.GET("/status", controller.Status)
	router.POST("/status", controller.Status)
	router.GET("/status/:statusCode", controller.Status)
	router.POST("/status/:statusCode", controller.Status)

	router.GET("/rand_status", controller.RandResult)

	router.GET("/version", controller.Version)

	router.GET("/get", controller.GetHandler)

	router.GET("/api/demo", controller.ApiDemo)
	router.GET("/api/crossservicecall", controller.CrossServiceCall)

	if runPort == "" {
		runPort = "3000"
	}
	router.Run(":" + runPort)
}
