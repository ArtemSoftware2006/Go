package main

import (
	"log"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	// Закрываем файл после выхода из main
	defer file.Close()
	// Конфигурируем логгер, чтобы он выводил лог в файл
	log.SetOutput(file)
	log.Println(message)

	return nil
}
