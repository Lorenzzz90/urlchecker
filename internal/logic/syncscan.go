package logic

import (
	"Lorenzzz90/urlchecker/tools"
	"fmt"
	"net/http"
	"time"
)

func SyncScan(urls []string) {
	for s := range urls {
		resp, err := http.Get(urls[s])
		tools.Check(err)
		defer resp.Body.Close()
		fmt.Printf("%s %s: Response %s\n", time.Now(), urls[s], resp.Status)
	}
}
