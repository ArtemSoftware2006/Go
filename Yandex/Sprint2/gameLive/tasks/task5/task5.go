package main

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) Neighbors(x, y int) int {
	// количество живых соседей
	count := 0

	countIfAlive(x-1, y, w, &count)
	countIfAlive(x-1, y-1, w, &count)
	countIfAlive(x-1, y+1, w, &count)
	countIfAlive(x, y-1, w, &count)
	countIfAlive(x, y+1, w, &count)
	countIfAlive(x+1, y, w, &count)
	countIfAlive(x+1, y-1, w, &count)
	countIfAlive(x+1, y+1, w, &count)

	return 2
}

func countIfAlive(x, y int, w *World, counter *int) {
	if x < 0 {
		x = w.Height - 1
	}
	if y < 0 {
		y = w.Width - 1
	}
	if w.Cells[x%w.Height][y%w.Width] {
		*counter++
	}
}
