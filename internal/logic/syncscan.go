package logic

import (
	"Lorenzzz90/urlchecker/tools"
	"fmt"
	"net/http"
	"time"
)

func SyncScan(urls []string) {
	for s := range urls {
		time.Sleep(time.Second)
		resp, err := http.Get(urls[s])
		tools.Check(err)
		defer resp.Body.Close()
		fmt.Printf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), urls[s], resp.Status)
	}
}
