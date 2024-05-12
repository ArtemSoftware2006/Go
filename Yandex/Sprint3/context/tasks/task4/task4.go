package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	// Создаем маршрутизатор HTTP
	mux := http.NewServeMux()

	// Обработчик для /readSource
	mux.HandleFunc("/readSource", func(w http.ResponseWriter, r *http.Request) {
		// Создаем контекст с таймаутом
		ctx, cancel := context.WithTimeout(r.Context(), maxTimeout)
		defer cancel()

		// Создаем запрос к серверу provideData
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8081/provideData", nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Создаем клиент с таймаутом
		client := &http.Client{
			Timeout: maxTimeout,
		}

		// Делаем запрос к серверу provideData
		resp, err := client.Do(req)
		if err != nil {
			// Проверяем, если ошибка произошла из-за таймаута
			if ctx.Err() == context.DeadlineExceeded {
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
		defer resp.Body.Close()

		// Читаем тело ответа и отправляем его обратно клиенту
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// Запускаем веб-сервер на порту 8080
	http.ListenAndServe(":8080", http.TimeoutHandler(mux, maxTimeout, "Service Unavailable"))
}
