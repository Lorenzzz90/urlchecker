package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"Lorenzzz90/urlchecker/internal/logic"
	"Lorenzzz90/urlchecker/tools"
)

// set global variables for args
var (
	consoleOutput *bool = flag.Bool("c", false, "prints output in the console instead of on a file")
	multipleFiles *bool = flag.Bool("m", false, "creates a new file for every single url, default: false")
	syncyes       *bool = flag.Bool("s", false, "decides if the program should run in sync or async mode, default: false")
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
	defer func() { fmt.Println(time.Since(start)) }()

	flag.Parse()
	// TODO: use meaningful names. Define a custom types (e.g. an Enum in C#)
	// it's hard to guess the meaning of these letters.
	var outputMode byte
	switch {
	case *consoleOutput:
		outputMode = 'c'
	case *multipleFiles:
		outputMode = 'm'
	default:
		outputMode = 'd'
	}
	// open the file containing the list of urls and append them to the list urls[]
	readFile, err := os.Open("./urls.txt")
	tools.Check(err)
	defer readFile.Close()

	// FIXME: manage the blank lines within the "urls.txt" file anywhere.
	// TODO: write output files to a "/tmp" folder or something similar (this folder should be added to the .gitignore file, otherwise each time you change something the working directory will be outdated.)
	scanner := bufio.NewScanner(readFile)
	var urls []string

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	tools.Check(scanner.Err())
	if *syncyes {
		logic.SyncScan(urls, outputMode)
	} else {
		logic.AsyncScan(urls, outputMode)
	}
}
