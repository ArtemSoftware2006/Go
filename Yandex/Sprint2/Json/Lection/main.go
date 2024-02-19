package main

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	decode()
}
