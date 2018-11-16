package utils

import (
	"os"
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
