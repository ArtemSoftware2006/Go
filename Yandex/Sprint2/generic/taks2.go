package main

type Stack[T any] struct {
	data []T
}

func (stack *Stack[T]) Push(val T) {
	stack.data = append(stack.data, val)
}

func (stack *Stack[T]) Pop() T {
	if len(stack.data) == 0 {
		panic("Stack is empty")
	}
	val := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return val
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}
