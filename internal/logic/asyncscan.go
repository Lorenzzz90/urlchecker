package logic

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"Lorenzzz90/urlchecker/tools"
)

func AsyncScan(urls []string, outputMode string) error {
	done := make(chan bool, len(urls))
	var writeFile *os.File
	var err error
	// FIXME: you used an "if", "elseif", but the "else"?
	// if the output mode is neither 'd' or 'm' there's no need to create a file or a folder so i left the else, is it always wrong to not have an else statement?
	if outputMode == "Default" {
		// FIXME: you MUST defer the close func on the file handle.
		writeFile, err = os.Create("tmp/" + tools.Today() + ".txt")
		if err != nil {
			return fmt.Errorf("err creating folder: %w", err)
		}
		defer writeFile.Close()
	} else if outputMode == "WriteToMultipleFiles" {
		if _, err := os.Stat("tmp" + tools.Today()); os.IsNotExist(err) {
			err := os.Mkdir("tmp/"+tools.Today(), os.ModeAppend)
			if err != nil {
				return fmt.Errorf("err creating folder: %w", err)
			}
		}
	}
	for i := 0; i < len(urls); i++ {
		go scanUrl(done, urls[i], outputMode, writeFile)
	}
	for i := 0; i < len(urls); i++ {
		<-done
	}
	return nil
}

func scanUrl(done chan bool, url string, outputMode string, writeFile *os.File) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("err getting url: %w", err)
	}
	if outputMode == "WriteToConsole" {
		fmt.Printf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status)
	} else if outputMode == "Default" {
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
	} else if outputMode == "WriteToMultipleFiles" {
		writeFile, err := os.Create(tools.CreatePath(url))
		writeFile.WriteString(fmt.Sprintf("tmp/%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), url, resp.Status))
		if err != nil {
			return fmt.Errorf("err writing to file: %w", err)
		}
		defer writeFile.Close()
	}
	defer resp.Body.Close()
	done <- true
	return nil
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
