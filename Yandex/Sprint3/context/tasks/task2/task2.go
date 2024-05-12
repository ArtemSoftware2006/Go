package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string // тело ответа
	StatusCode int    // код ответа
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ctxWithTime, _ := context.WithTimeout(ctx, timeout)

	resp, err := client.Do(req.WithContext(ctxWithTime))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, context.DeadlineExceeded
	default:
		return &APIResponse{
			Data:       string(body),
			StatusCode: resp.StatusCode,
		}, nil
	}
}
