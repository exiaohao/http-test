package utils

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

// Hostname	current pod or server name
func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

// Version current env:Version
func Version() string {
	userVersion := os.Getenv("VERSION")
	if userVersion == "" {
		userVersion = "v1.0/default-version"
	}
	return userVersion
}

// RandomHTTPStatus
func RandomHTTPStatus() (int, int) {
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
	return errRate, statusCode
}
