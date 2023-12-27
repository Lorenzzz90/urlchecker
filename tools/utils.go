package tools

import (
	"fmt"
	"strings"
	"time"
)

func CreatePath(url string) string {
	today := Today()
	fileName := fmt.Sprintf("%s_%s", today, url) // TODO usare regexp per bonificare url e usare package filepath
	fileName = strings.Replace(fileName, "https://", "", -1)
	fileName = strings.Replace(fileName, ".", "_", -1)
	fileName = strings.Replace(fileName, "/", "-", -1)
	path := fmt.Sprintf("tmp/%s/%s.txt", today, fileName)
	return path
}

func Today() string {
	return time.Now().Format("2006-01-02")
}
