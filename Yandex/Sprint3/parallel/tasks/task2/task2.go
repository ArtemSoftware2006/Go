package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// WordCounter структура для подсчета количества вхождений слов во всех файлах.
type WordCounter struct {
	wordsCount map[string]int
	mu         sync.Mutex
}

// CounterWorker интерфейс для обработки файлов и подсчета слов.
type CounterWorker interface {
	ProcessFiles(files ...string) error
	ProcessReader(r io.Reader) error
}

// NewWordCounter создает новый экземпляр WordCounter.
func NewWordCounter() *WordCounter {
	return &WordCounter{
		wordsCount: make(map[string]int),
	}
}

// ProcessFiles запускает обработку файлов.
func (wc *WordCounter) ProcessFiles(files ...string) error {
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := wc.processFile(filename)
			if err != nil {
				fmt.Printf("Ошибка при обработке файла %s: %v\n", filename, err)
			}
		}(file)
	}
	wg.Wait()
	return nil
}

// ProcessReader подсчитывает слова в одном файле.
func (wc *WordCounter) ProcessReader(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text()) // Преобразование слова в нижний регистр
		wc.mu.Lock()
		wc.wordsCount[word]++
		wc.mu.Unlock()
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// processFile открывает файл, считывает его и передает на обработку в ProcessReader.
func (wc *WordCounter) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := wc.ProcessReader(file); err != nil {
		return err
	}
	return nil
}
