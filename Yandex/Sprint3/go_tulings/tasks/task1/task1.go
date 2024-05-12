package task1

import "sync"

func RaceConditionFunc(counter *int, wg *sync.WaitGroup) {
	var mu sync.Mutex

	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}
