package logic

import (
	"fmt"
	"net/http"
	"time"
)

func AsyncScan(urls []string) ([]string, error) {
	var responses []string
	respChannel := make(chan string, len(urls))
	done := make(chan bool, len(urls))
	for url := range urls {
		go scanUrl(urls[url], done, respChannel)
	}
	for i := 0; i < len(urls); i++ {
		responses = append(responses, <-respChannel)
		<-done
	}
	return responses, nil
}

func scanUrl(url string, done chan bool, respChannel chan string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("err getting url: %w", err)
	}
	defer resp.Body.Close()
	respChannel <- fmt.Sprintf("%s %s: Response %s", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status)
	done <- true
	return nil
}
