package task2

import "sync"

func SafeCounterFunc(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}
