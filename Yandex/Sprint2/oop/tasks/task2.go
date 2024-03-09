package main

type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string, age int) *User {
	return &User{
		Name:   name,
		Age:    age,
		Active: true,
	}
}
