package controller

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/exiaohao/http-test/utils"
	"github.com/gin-gonic/gin"
)

var randStatuses = []int{
	http.StatusOK,
	http.StatusNoContent,
	http.StatusBadRequest,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusFound,
	http.StatusNotFound,
	http.StatusTeapot,
}

// Status return httpStatus what you want
func Status(c *gin.Context) {
	statusCode, _ := strconv.Atoi(c.Param("statusCode"))

	if statusCode == 0 {
		idx := rand.Intn(len(randStatuses))
		statusCode = randStatuses[idx]
	}

	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"statusText": http.StatusText(statusCode),
		"serverName": utils.Hostname(),
		"version":    utils.Version(),
	})
}

// RandResult returns random result
func RandResult(c *gin.Context) {
	var errRate, statusCode int
	errRateString := os.Getenv("ERR_RATE")
	if errRateString == "" {
		errRate = 50
	} else {
		errRate, _ = strconv.Atoi(errRateString)
	}

	if rand.Intn(100) < errRate {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"statusText": http.StatusText(statusCode),
		"serverName": utils.Hostname(),
		"version":    utils.Version(),
	})
}

// Version diy your version
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"version":    utils.Version(),
		"serverName": utils.Hostname(),
	})
}

// GetHandler response get request
func GetHandler(c *gin.Context) {
	getRequest := make(gin.H)
	getHeaders := make(gin.H)

	for key, val := range c.Request.Header {
		getHeaders[key] = val[0]
	}

	for key, val := range c.Request.URL.Query() {
		getRequest[key] = val[0]
	}

	c.JSON(http.StatusOK, gin.H{
		"args":       getRequest,
		"headers":    getHeaders,
		"origin":     c.Request.RemoteAddr,
		"statusCode": http.StatusOK,
		"serverName": utils.Hostname(),
		"version":    utils.Version(),
	})
}
