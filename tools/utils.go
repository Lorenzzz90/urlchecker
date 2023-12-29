package tools

import (
	"strings"
	"time"
)

func CreateFileName(response string) string {
	parts := strings.Split(response, "www")
	fileName := parts[1]
	fileName = strings.Replace(fileName, "/", "", -1)
	fileName = strings.Replace(fileName, ".", "-", -1)
	fileName = strings.Replace(fileName, " ", "", -1)
	fileName = strings.Replace(fileName, ":", "", -1)
	parts2 := strings.Split(fileName, "Response")
	fileName = parts2[0]
	return fileName
}

func Today() string {
	return time.Now().Format("2006-01-02")
}
