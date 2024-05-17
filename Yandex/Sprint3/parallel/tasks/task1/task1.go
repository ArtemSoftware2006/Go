package main

import (
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	out := make(chan Work)

	go func() {
		defer close(out) // закроем канал после обхода всех файлов
		// функция для перебора файлов в директории
		filepath.Walk(dir, func(path string, d fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// пропускаем вложенные директории
			if !d.IsDir() {
				// запишем в канал файл, которые нужно обработать на следующем этапе
				out <- Work{FilePath: path}
			}
			return nil
		})
	}()

	return out
}
func compress(jobs <-chan Work) {

	var wg sync.WaitGroup

	for work := range jobs {

		wg.Add(1)

		go func(work Work) {
			defer wg.Done()
			inputFile, err := os.Open(work.FilePath)
			if err != nil {
				return
			}
			defer inputFile.Close()

			outputPath := work.FilePath + ".gz"
			outputFile, err := os.Create(outputPath)
			if err != nil {
				return
			}
			defer outputFile.Close()

			gzipWriter := gzip.NewWriter(outputFile)
			defer gzipWriter.Close()

			_, err = io.Copy(gzipWriter, inputFile)
			if err != nil {
				return
			}

		}(work)
	}

	wg.Wait()
}

type Work struct {
	FilePath string
}
