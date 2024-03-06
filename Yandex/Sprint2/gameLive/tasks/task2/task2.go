package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func main() {
	w := World{}
	err := w.LoadState("state.bin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("w: %v\n", w)

}

func (w *World) LoadState(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			return errors.New("empty name")
		}

		w.Height++

		if w.Width == 0 {
			w.Width = len(line)
		}
		if w.Width != len(line) {
			return errors.New("wrong width")
		}

		w.Cells = append(w.Cells, []bool{})

		for i := 0; i < w.Width; i++ {
			c := line[i]
			if c == '1' {
				w.Cells[w.Height-1] = append(w.Cells[w.Height-1], true)
			} else {
				w.Cells[w.Height-1] = append(w.Cells[w.Height-1], false)
			}
		}
	}

	return nil
}
