package logic

import (
	"Lorenzzz90/urlchecker/tools"
	"fmt"
	"net/http"
	"os"
	"time"
)

func SyncScan(urls []string, multipleFiles bool) {
	if !multipleFiles {
		writeFile, err := os.Create(time.Now().Format("2006-01-02") + ".txt")
		tools.Check(err)
		defer writeFile.Close()
		for s := range urls {
			resp, err := http.Get(urls[s])
			tools.Check(err)
			defer resp.Body.Close()
			writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", urls[s], resp.Status))
		}
	} else {
		if _, err := os.Stat(time.Now().Format("2006-01-02")); os.IsNotExist(err) {
			err := os.Mkdir(time.Now().Format("2006-01-02"), os.ModeAppend)
			tools.Check(err)
		}
		for s := range urls {
			writeFile, err := os.Create(tools.CreatePath(urls[s]))
			tools.Check(err)
			resp, err := http.Get(urls[s])
			tools.Check(err)
			defer resp.Body.Close()
			writeFile.WriteString(fmt.Sprintf("%s: Response %s\n", urls[s], resp.Status))
		}
	}
}
