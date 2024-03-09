package main

import (
	"fmt"
	"time"
)

func main() {
	// Устанавливаем часовой пояс UTC
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Ошибка при загрузке часового пояса:", err)
		return
	}

	// Получаем текущее время в указанном часовом поясе
	currentTime := time.Now().In(loc)

	// Форматируем время в строку с заданным форматом
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Println("Текущее время (UTC):", formattedTime)
}
