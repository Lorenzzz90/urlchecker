package tools

import (
	"strings"
	"time"
)

// FIXME: IMO is better to move this pkg below internal as well.

func CreateFileName(response string) string {
	// TODO: this could be rewritten with a regexp to make it nicer.
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
