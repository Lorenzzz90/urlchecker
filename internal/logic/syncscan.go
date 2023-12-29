package logic

import (
	"fmt"
	"net/http"
	"time"
)

func SyncScan(urls []string) ([]string, error) {
	var responses []string

	for s := range urls {
		resp, err := http.Get(urls[s])
		if err != nil {
			return nil, fmt.Errorf("err getting url: %w", err)
		}
		defer resp.Body.Close()
		msg := fmt.Sprintf("%s %s: Response %s", time.Now().Format("2006-01-02T15:04:05"), urls[s], resp.Status)
		responses = append(responses, msg)
	}
	return responses, nil
}
