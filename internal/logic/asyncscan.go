package logic

import (
	"Lorenzzz90/urlchecker/tools"
	"fmt"
	"net/http"
	"time"
)

func AsyncScan(urls []string) {
	done := make(chan bool, len(urls))
	for i := 0; i < len(urls); i++ {
		go scanUrl(done, urls[i])
	}
	for i := 0; i < len(urls); i++ {
		<-done
	}

}

func scanUrl(done chan bool, url string) {
	time.Sleep(time.Second)
	resp, err := http.Get(url)
	tools.Check(err)
	fmt.Printf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status)
	defer resp.Body.Close()
	done <- true
}

/*func AsyncScan(urls []string, multipleFiles bool) {
	var writeFile *os.File
	var err error
	if !multipleFiles {
		writeFile, err = os.Create(tools.Today() + ".txt")
		tools.Check(err)
	} else {
		if _, err := os.Stat(tools.Today()); os.IsNotExist(err) {
			err := os.Mkdir(tools.Today(), os.ModeAppend)
			tools.Check(err)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			scanUrl(urls[i], tools.CreatePath(urls[i]), writeFile, multipleFiles)
		}()
	}
	wg.Wait()
}

// function called by goroutine insine asyncScan function
func scanUrl(url string, path string, writeFile *os.File, multipleFiles bool) {
	resp, err := http.Get(url)
	tools.Check(err)
	if !multipleFiles {
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
	} else {
		writeFile, err := os.Create(tools.CreatePath(url))
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
		tools.Check(err)
	}

	defer resp.Body.Close()

}*/
