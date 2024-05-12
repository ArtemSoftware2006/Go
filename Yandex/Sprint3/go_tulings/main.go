package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Создание файла для записи дампа горутин
	f, err := os.Create("goroutine_dump.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer f.Close()

	// Запись дампа горутин в файл каждую секунду
	go func() {
		for {
			// Получение дампа горутин
			pprof.Lookup("goroutine").WriteTo(f, 1)
			time.Sleep(time.Second)
		}
	}()

	// Ваш код, в котором создаются горутины
	// ...

	// Пример бесконечного цикла для демонстрации
	for {
		// Здесь может быть ваша основная логика программы
		time.Sleep(1 * time.Second)
	}
}
