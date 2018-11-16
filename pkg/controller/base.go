package controller

import (
	"math/rand"
	"net/http"
	"strconv"

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
	})
}
