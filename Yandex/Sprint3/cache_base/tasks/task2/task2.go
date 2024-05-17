package main

import (
	"container/list"
	"sync"
)

// структура MRU кеша
type MRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

type CacheEntry struct {
	key   string
	value string
}

// возвращает новый инстанс кеша размером capacity
func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		mutex:    sync.Mutex{},
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// устанавливает значени value ключу key
func (c *MRUCache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		element.Value.(*CacheEntry).value = value
		c.list.MoveToFront(element)
	} else {
		entry := &CacheEntry{key: key, value: value}
		element = c.list.PushFront(entry)
		c.cache[key] = element

		if c.capacity < c.list.Len() {
			youngest := c.list.Back()
			if youngest != nil {
				delete(c.cache, youngest.Value.(*CacheEntry).key)
				c.list.Remove(youngest)
			}
		}
	}
}

// получает значение и флаг его начличия по ключу key
func (c *MRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		c.list.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}

	return "", false
}
