package task1

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	file, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counterLines := 0
	for scanner.Scan() {
		if counterLines == lineNum {
			return scanner.Text()
		}
		counterLines++
	}

	return ""
}
