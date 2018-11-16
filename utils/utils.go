package utils

import (
	"os"
)

func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
