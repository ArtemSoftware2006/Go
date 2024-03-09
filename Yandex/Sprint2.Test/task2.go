package main

import "errors"

type Stack[T any] struct {
	data []T
}

func (stack *Stack[T]) Push(key T) {
	stack.data = append(stack.data, key)
}

func (stack *Stack[T]) Pop() (T, error) {
	if len(stack.data) == 0 {
		return *new(T), errors.New("Stack is empty")
	}

	val := stack.data[len(stack.data)-1]
	stack.data = append(stack.data[0 : len(stack.data)-1])

	return val, nil
}
