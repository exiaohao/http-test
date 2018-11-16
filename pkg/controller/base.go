package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/exiaohao/http-test/utils"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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

// ApiDemoReturn
type ApiDemoReturn struct {
	Data       string `json:"data"`
	StatusCode int    `json:"statusCode"`
	ServerName string `json:"serverName"`
	Version    string `json:"version"`
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
	_, statusCode := utils.RandomHTTPStatus()

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

// ApiDemo a api demo
func ApiDemo(c *gin.Context) {
	_, statusCode := utils.RandomHTTPStatus()

	randData := uuid.Must(uuid.NewV4())
	c.JSON(statusCode, gin.H{
		"data":       randData,
		"statusCode": http.StatusOK,
		"serverName": utils.Hostname(),
		"version":    utils.Version(),
	})
}

// CrossServiceCall
func CrossServiceCall(c *gin.Context) {
	requestUrl := os.Getenv("TARGET_SERVICE") + "/api/demo"
	resp, err := http.Get(requestUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"err":        fmt.Sprintf("Request backend %s failed, because: %s", requestUrl, err),
			"statusCode": http.StatusInternalServerError,
			"serverName": utils.Hostname(),
			"version":    utils.Version(),
		})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"err":        fmt.Sprintf("Read body failed, because: %s", err),
			"statusCode": http.StatusInternalServerError,
			"serverName": utils.Hostname(),
			"version":    utils.Version(),
		})
		return
	}

	returnBodyJson := ApiDemoReturn{}
	err = json.Unmarshal(body, &returnBodyJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"rawData":    string(body),
			"err":        fmt.Sprintf("Parse json failed, because: %s", err),
			"statusCode": http.StatusInternalServerError,
			"serverName": utils.Hostname(),
			"version":    utils.Version(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"data":       returnBodyJson.Data,
			"statusCode": returnBodyJson.StatusCode,
			"serverName": returnBodyJson.ServerName,
			"version":    returnBodyJson.Version,
		},
		"err":        nil,
		"statusCode": http.StatusOK,
		"serverName": utils.Hostname(),
		"version":    utils.Version(),
	})
}
