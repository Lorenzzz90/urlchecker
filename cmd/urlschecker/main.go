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

var (
	WriteToConsole       *bool = flag.Bool("c", false, "prints output in the console instead of on a file")
	WriteToMultipleFiles *bool = flag.Bool("m", false, "creates a new file for every single url, default: false")
	useSyncMod           *bool = flag.Bool("s", false, "decides if the program should run in sync or async mode, default: false")
)

func main() {
	start := time.Now()
	defer func() { fmt.Printf("Execution Time: %v", time.Since(start)) }()

	flag.Parse()

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

	var responses []string
	if *useSyncMod {
		responses, err = logic.SyncScan(urls)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		responses, err = logic.AsyncScan(urls)
		if err != nil {
			fmt.Println(err)
		}
	}

	switch {
	case *WriteToConsole:
		logic.PrintToConsole(responses)
	case *WriteToMultipleFiles:
		logic.PrintToMultipleFiles(responses)
	default:
		logic.PrintToFile(responses)
	}
}
