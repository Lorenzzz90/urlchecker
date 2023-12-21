package main

import (
	"Lorenzzz90/urlchecker/internal/logic"
	"Lorenzzz90/urlchecker/tools"
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

// set global variables for args
var consoleOutput *bool = flag.Bool("c", false, "prints output in the console instead of on a file")
var multipleFiles *bool = flag.Bool("m", false, "creates a new file for every single url, default: false")
var syncyes *bool = flag.Bool("s", false, "decides if the program should run in sync or async mode, default: false")

func main() {
	//output execution time
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	flag.Parse()
	var outputMode byte
	switch {
	case *consoleOutput:
		outputMode = 'c'
	case *multipleFiles:
		outputMode = 'm'
	default:
		outputMode = 'd'
	}
	//open the file containing the list of urls and append them to the list urls[]
	readFile, err := os.Open("./urls.txt")
	tools.Check(err)
	defer readFile.Close()

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
