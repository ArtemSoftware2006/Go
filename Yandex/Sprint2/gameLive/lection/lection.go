package main

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func main() {
	// зададим размеры сетки
	height := 10
	width := 10
	// объект для хранения текущего состояния сетки
	currentWorld := NewWorld(height, width)
	// объект для хранения следующего состояния сетки
	nextWorld := NewWorld(height, width)
	// установим начальное состояние
	currentWorld.Seed()
	for { // цикл для вывода каждого состояния
		// выведем текущее состояние на экран
		fmt.Println(currentWorld)
		// рассчитываем следующее состояние
		NextState(currentWorld, nextWorld)
		// изменяем текущее состояние
		currentWorld = nextWorld
		// делаем паузу
		time.Sleep(100 * time.Millisecond)
		// специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
	}
}

func (w *World) Seed() {
	// снова переберём все клетки
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func NextState(oldWorld, newWorld *World) {
	// переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) Neighbours(x, y int) int {
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

	return count
}

func countIfAlive(x, y int, w *World, counter *int) {
	if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
		return
	} else if w.Cells[x][y] {
		*counter++
	}
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)      // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
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
