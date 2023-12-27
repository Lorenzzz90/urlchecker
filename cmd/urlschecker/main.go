package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"Lorenzzz90/urlchecker/internal/logic"
)

// set global variables for args
var (
	WriteToConsole       *bool = flag.Bool("c", false, "prints output in the console instead of on a file")
	WriteToMultipleFiles *bool = flag.Bool("m", false, "creates a new file for every single url, default: false")
	useSyncMod           *bool = flag.Bool("s", false, "decides if the program should run in sync or async mode, default: false")
)

// TODO: depending on where you're located in the file system, it can give you an error such as the following one:
/*
go run cmd/urlschecker/main.go -c
70.664µs
panic: open ./urls.txt: no such file or directory

goroutine 1 [running]:
Lorenzzz90/urlchecker/tools.Check(...)
        /media/ivan/Volume/training/lorenzo-marziali/urlchecker/tools/utils.go:21
main.main()
        /media/ivan/Volume/training/lorenzo-marziali/urlchecker/cmd/urlschecker/main.go:35 +0x338
exit status 2
*/
func main() {
	// output execution time
	start := time.Now()
	defer func() { fmt.Printf("Execution Time: %v", time.Since(start)) }()

	flag.Parse()
	// TODO: use meaningful names. Define a custom types (e.g. an Enum in C#)
	// it's hard to guess the meaning of these letters.
	var outputMode string
	switch {
	case *WriteToConsole:
		outputMode = "WriteToConsole" //Writes the output of the files to the console
	case *WriteToMultipleFiles:
		outputMode = "WriteToMultipleFiles" //Writes the output in a folder in which a single file is created for each url
	default:
		outputMode = "Default" //Writes the output to a single txt file containing all the urls
	}
	// open the file containing the list of urls and append them to the list urls[]
	//filepath, err := filepath.Abs("./urls.txt")
	//tools.Check(err)
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	readFile, err := os.Open(basepath + "./urls.txt")
	if err != nil {
		fmt.Println("Error reading file urls.txt", err.Error())
		return
	}
	defer readFile.Close()

	// FIXME: manage the blank lines within the "urls.txt" file anywhere.
	// TODO: write output files to a "/tmp" folder or something similar (this folder should be added to the .gitignore file, otherwise each time you change something the working directory will be outdated.)
	scanner := bufio.NewScanner(readFile)
	var urls []string

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		} else {
			urls = append(urls, scanner.Text())
		}
	}
	if scanner.Err() != nil {
		fmt.Println("error scanning file", scanner.Err().Error())
		return
	}
	if outputMode != "WriteToConsole" {
		if _, err := os.Stat("tmp"); os.IsNotExist(err) {
			err := os.Mkdir("tmp", os.ModeAppend)
			if err != nil {
				fmt.Println("error creating folder tmp", err.Error())
			}
		}
	}
	if *useSyncMod {
		logic.SyncScan(urls, outputMode)
	} else {
		logic.AsyncScan(urls, outputMode)
	}
}
