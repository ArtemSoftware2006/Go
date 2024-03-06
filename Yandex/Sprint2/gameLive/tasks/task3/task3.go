package main

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"

	outputString := ""

	for i := 0; i < len(w.Cells); i++ {
		for j := 0; j < len(w.Cells[i]); j++ {
			if w.Cells[i][j] {
				outputString += greenSquare
			} else {
				outputString += brownSquare
			}
		}
		outputString += "\n"
	}

	return outputString
}
