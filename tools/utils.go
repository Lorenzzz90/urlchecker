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
	path := fmt.Sprintf("%s/%s.txt", today, fileName)
	return path
}

func Check(e error) {
	if e != nil {
		// FIXME: "panic" is hard to recover for a program. You need to use recover()
		// instead bubble up the error by wrapping it with the func fmt.Errof and the %w verb.
		// plus, this func can be avoided at all and this logic could be spread among the functions and/or methods.
		panic(e)
	}
}

func Today() string {
	return time.Now().Format("2006-01-02")
}
