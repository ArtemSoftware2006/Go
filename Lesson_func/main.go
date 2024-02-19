package main

import (
	"fmt"
)

func main() {

	var function func(int) int = sqr
	function_2 := func(x int) int { return x * x } //Анонимная функция

	fmt.Println(function(5))

	for i := 1; i < 10; i++ {
		//fmt.Println(sqr(i))

		var value, str = magicSqr(i)

		fmt.Println(value, str)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			//panic("Error: i is 3") // panic - генерация ошибки
		}
		fmt.Println(function_2(i))
	}

	defer finish() // defer - выполняется в самом конце функции
	defer fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func sqr(a int) int {
	return a * a
}

func magicSqr(a int) (int, string) {
	return a * a, string(a) + " is magic"
}

func finish() {
	fmt.Println("Program has been finished")
}
