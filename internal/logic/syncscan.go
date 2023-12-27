package logic

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"Lorenzzz90/urlchecker/tools"
)

func SyncScan(urls []string, outputMode string) error {
	// FIXME: lot of duplicated logic.
	// the only thing that changes here is the output destination. Try to shorten this code.
	for s := range urls {
		resp, err := http.Get(urls[s])
		if err != nil {
			return fmt.Errorf("err getting url: %w", err)
		}
		defer resp.Body.Close()
		if outputMode == "WriteToConsole" {
			fmt.Printf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), urls[s], resp.Status)
		} else if outputMode == "Default" {
			var writeFile *os.File
			if _, err := os.Stat("tmp/" + tools.Today() + ".txt"); os.IsNotExist(err) {
				writeFile, err = os.Create("tmp/" + tools.Today() + ".txt")
				if err != nil {
					return fmt.Errorf("err creating tmp directory: %w", err)
				}
			}
			defer writeFile.Close()
			writeFile.WriteString(fmt.Sprintf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), urls[s], resp.Status))
		} else if outputMode == "WriteToMultipleFiles" {
			if _, err := os.Stat("tmp" + tools.Today()); os.IsNotExist(err) {
				err := os.Mkdir("tmp/"+tools.Today(), os.ModeAppend)
				if err != nil {
					return fmt.Errorf("err creating today folder: %w", err)
				}
			}
			if _, err := os.Stat("tmp/" + time.Now().Format("2006-01-02")); os.IsNotExist(err) {
				err := os.Mkdir("tmp/"+time.Now().Format("2006-01-02"), os.ModeAppend)
				if err != nil {
					return fmt.Errorf("error creating today folder%w", err)
				}
			}
			writeFile, err := os.Create(tools.CreatePath(urls[s]))
			if err != nil {
				return fmt.Errorf("error writing to file: %w", err)
			}
			defer writeFile.Close()
			writeFile.WriteString(fmt.Sprintf("%s %s: Response %s\n", time.Now().Format("2006-01-02T15:04:05"), urls[s], resp.Status))
		}
	}
	return nil
}
