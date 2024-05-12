package task1

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	return s.m[key]
}
func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}
