package main

import (
	"testing"
)

func TestMRUCache(t *testing.T) {
	cache := NewMRUCache(2)

	// Тест добавления элементов в кэш
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Проверка добавленных значений
	value, ok := cache.Get("key1")
	if !ok || value != "value1" {
		t.Errorf("Expected value for key1 to be 'value1', but got '%s'", value)
	}

	value, ok = cache.Get("key2")
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}

	// Обновление значения элемента в кэше
	cache.Set("key1", "new value1")
	value, ok = cache.Get("key1")
	if !ok || value != "new value1" {
		t.Errorf("Expected value for key1 to be 'new value1', but got '%s'", value)
	}

	// Проверка вытеснения наименее недавно используемого элемента (MRU логика)
	cache.Set("key3", "value3") // This should evict "key2" since it's MRU
	_, ok = cache.Get("key2")
	if ok {
		t.Error("Expected 'key2' to be evicted from the cache")
	}

	value, ok = cache.Get("key3")
	if !ok || value != "value3" {
		t.Errorf("Expected value for key3 to be 'value3', but got '%s'", value)
	}

	value, ok = cache.Get("key1")
	if !ok || value != "new value1" {
		t.Errorf("Expected value for key1 to be 'new value1', but got '%s'", value)
	}

	// Проверка, что кэш обновляется правильно, и самые недавно использованные элементы остаются
	cache.Set("key4", "value4") // This should evict "key1" now
	_, ok = cache.Get("key1")
	if ok {
		t.Error("Expected 'key1' to be evicted from the cache")
	}

	value, ok = cache.Get("key3")
	if !ok || value != "value3" {
		t.Errorf("Expected value for key3 to be 'value3', but got '%s'", value)
	}

	value, ok = cache.Get("key4")
	if !ok || value != "value4" {
		t.Errorf("Expected value for key4 to be 'value4', but got '%s'", value)
	}
}
