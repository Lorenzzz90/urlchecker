package logic

import (
	"fmt"
	"os"

	"Lorenzzz90/urlchecker/tools"
)

func PrintToConsole(responses []string) {
	for i := range responses {
		fmt.Println(responses[i])
	}
}

func PrintToFile(responses []string) error {
	writeFile, err := os.Create("../../tmp/" + tools.Today() + ".txt")
	if err != nil {
		return fmt.Errorf("err creating txt file: %w", err)
	}
	defer writeFile.Close()
	for i := range responses {
		// FIXME: check for the err
		writeFile.WriteString(responses[i] + "\n")
	}
	return nil
}

func PrintToMultipleFiles(responses []string) error {
	if _, err := os.Stat("../../tmp/" + tools.Today()); os.IsNotExist(err) {
		err := os.Mkdir("../../tmp/"+tools.Today(), os.ModeAppend)
		if err != nil {
			return fmt.Errorf("err creating today folder: %w", err)
		}
	}
	for i := range responses {
		writeFile, err := os.Create("../../tmp/" + tools.Today() + "/" + tools.CreateFileName(responses[i]) + ".txt")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
		defer writeFile.Close()
		writeFile.WriteString(fmt.Sprint(responses[i]))
	}
	return nil
}
