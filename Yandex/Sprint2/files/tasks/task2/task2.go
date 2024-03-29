package task2

import (
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(int64(pos), 0)
	file.WriteString(val)
}
