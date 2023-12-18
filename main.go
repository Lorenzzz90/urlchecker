package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// set global variables for args
var multipleFiles *bool = flag.Bool("m", false, "creates a new file for every single url, default: false")
var syncyes *bool = flag.Bool("s", false, "decides if the program should run in sync or async mode, default: false")

func main() {
	//output execution time
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	flag.Parse()
	//open the file containing the list of urls and append them to the list urls[]
	readFile, err := os.Open("./urls.txt")
	check(err)
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	var urls []string

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	check(scanner.Err())
	if *syncyes {
		syncScan(urls)
	} else {
		asyncScan(urls)
	}
}

// function called if the program is runned in sync mode
func syncScan(urls []string) {
	if !*multipleFiles {
		writeFile, err := os.Create(time.Now().Format("2006-01-02") + ".txt")
		check(err)
		for s := range urls {
			resp, err := http.Get(urls[s])
			check(err)
			writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", urls[s], resp.Status))
			defer resp.Body.Close()
		}
	} else {
		if _, err := os.Stat(time.Now().Format("2006-01-02")); os.IsNotExist(err) {
			err := os.Mkdir(time.Now().Format("2006-01-02"), os.ModeAppend)
			check(err)
		}
		for s := range urls {
			writeFile, err := os.Create(createPath(urls[s]))
			check(err)
			resp, err := http.Get(urls[s])
			check(err)
			writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", urls[s], resp.Status))
			defer resp.Body.Close()
		}
	}
}

// helper function to create the path of a file inside the main folder
func createPath(url string) string {
	today := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s_%s", today, url)
	fileName = strings.Replace(fileName, "https://", "", -1)
	fileName = strings.Replace(fileName, ".", "_", -1)
	fileName = strings.Replace(fileName, "/", "-", -1)
	path := fmt.Sprintf("%s/%s.txt", today, fileName)
	return path
}

// function called in the program is runned in async mode
func asyncScan(urls []string) {
	var writeFile *os.File
	var err error
	if !*multipleFiles {
		writeFile, err = os.Create(time.Now().Format("2006-01-02") + ".txt")
		check(err)
	} else {
		if _, err := os.Stat(time.Now().Format("2006-01-02")); os.IsNotExist(err) {
			err := os.Mkdir(time.Now().Format("2006-01-02"), os.ModeAppend)
			check(err)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			scanUrl(urls[i], createPath(urls[i]), writeFile)
		}()
	}
	wg.Wait()
}

// function called by goroutine insine asyncScan function
func scanUrl(url string, path string, writeFile *os.File) {
	resp, err := http.Get(url)
	check(err)
	if !*multipleFiles {
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
	} else {
		writeFile, err := os.Create(createPath(url))
		writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", url, resp.Status))
		check(err)
	}

	defer resp.Body.Close()

}

// helper function for checking for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}
