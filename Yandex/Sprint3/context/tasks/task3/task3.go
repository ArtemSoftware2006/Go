package task3

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var wg sync.WaitGroup
	responses := make([]*APIResponse, len(urls))
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			responses[i] = fetchAPISingle(ctx, url)

		}(i, url)
	}

	wg.Wait()
	return responses
}

func fetchAPISingle(ctx context.Context, url string) *APIResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &APIResponse{URL: url, Err: err}
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return &APIResponse{URL: url, Err: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIResponse{URL: url, Err: err}
	}

	select {
	case <-ctx.Done():
		return &APIResponse{URL: url, Err: context.DeadlineExceeded}
	default:
		return &APIResponse{
			URL:        url,
			Data:       string(body),
			StatusCode: resp.StatusCode,
		}
	}
}
