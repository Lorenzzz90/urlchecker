package logic

import (
	"Lorenzzz90/urlchecker/tools"
	"fmt"
	"net/http"
	"os"
	"time"
)

func AsyncScan(urls []string, outputMode byte) {
	done := make(chan bool, len(urls))
	var writeFile *os.File
	var err error
	if outputMode == 'd' {
		writeFile, err = os.Create(tools.Today() + ".txt")
		tools.Check(err)
	} else if outputMode == 'm' {
		if _, err := os.Stat(tools.Today()); os.IsNotExist(err) {
			err := os.Mkdir(tools.Today(), os.ModeAppend)
			tools.Check(err)
		}
	}
	for i := 0; i < len(urls); i++ {
		go scanUrl(done, urls[i], outputMode, writeFile)
	}
	for i := 0; i < len(urls); i++ {
		<-done
	}

}

func scanUrl(done chan bool, url string, outputMode byte, writeFile *os.File) {
	resp, err := http.Get(url)
	tools.Check(err)
	if outputMode == 'c' {
		fmt.Printf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status)
	} else if outputMode == 'd' {
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
	} else if outputMode == 'm' {
		writeFile, err := os.Create(tools.CreatePath(url))
		writeFile.WriteString(fmt.Sprintf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status))
		tools.Check(err)
		defer writeFile.Close()
	}
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
