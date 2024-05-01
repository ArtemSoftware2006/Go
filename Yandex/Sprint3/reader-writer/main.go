package main

import (
	"fmt"
	"io"
)

type myWriter struct {
	content []byte
}

func (w *myWriter) Write(buf []byte) (int, error) {
	w.content = append(w.content, buf...)
	return len(buf), nil
}

func (w *myWriter) String() string {
	return string(w.content)
}

func main() {
	w := &myWriter{}
	WriteData(w, []byte("Hello"))
	fmt.Println(string(w.content))
}

func WriteData(writer io.Writer, data []byte) error {
	bytesWritten, err := writer.Write(data) // Записываем данные
	if err != nil {
		// TODO: handle error.
	}
	fmt.Printf("Записано %d байт\n", bytesWritten)
	return nil
}
