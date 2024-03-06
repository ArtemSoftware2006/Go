package main

import (
	"fmt"
	"os"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func main() {
	w := NewWorld(3, 3)
	w.SaveState("state.bin")
}

func NewWorld(height, width int) *World {
	// создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if len(w.Cells) == 0 {
		return fmt.Errorf("empty world")
	}

	// Затем последовательно записываем значения ячеек сетки
	for i, row := range w.Cells {
		for _, cell := range row {
			var value string
			if cell {
				value = "1"
			} else {
				value = "0"
			}
			if _, err := file.WriteString(value); err != nil {
				return err
			}
		}
		if i != len(w.Cells)-1 {
			file.WriteString("\n")
		}
	}

	return nil
}
