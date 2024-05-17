package main

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func getProduct(productId int, db map[int]Product, cache *Cache) (Product, error) {
	if product, ok := cache.Get(productId); ok {
		return product, nil
	} else {
		product, ok := db[productId]
		if ok {
			cache.Set(productId, product)
			return product, nil
		} else {
			return Product{}, errors.New("Product not found")
		}
	}
}

//Функция обновления информации о товаре (фейк-функция выполняющая роль базы данных):

func updateProduct(productId int, newProduct Product, db map[int]Product) error {
	if _, ok := db[productId]; ok {
		db[productId] = newProduct
		return nil
	} else {
		db[productId] = newProduct
		return errors.New(fmt.Sprintf("Product not found (ProductId %d)", productId))
	}
}

//Кеш продуктов:

type Cache struct {
	products map[int]Product // Кэш продуктов
	ttl      time.Duration   // Время жизни записи в кэше
}

func NewCache() *Cache {
	return &Cache{
		products: make(map[int]Product),
		ttl:      time.Hour,
	}
}

//Получение продукта из кеша:

func (c *Cache) Get(productId int) (Product, bool) {
	if product, ok := c.products[productId]; ok {
		return product, true
	} else {
		return Product{}, false
	}
}

//Сохранение продукта в кеш:

func (c *Cache) Set(productId int, product Product) {
	c.products[productId] = product
}

//Инвалидация кеша:

func (c *Cache) Invalidate(productId int) {
	delete(c.products, productId)
}
