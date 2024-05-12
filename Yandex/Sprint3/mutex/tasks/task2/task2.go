package task2

import "sync"

type Counter struct {
	value int // значение счетчика
	mu    sync.RWMutex
}

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
