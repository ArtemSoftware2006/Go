package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//carta2 := make(map[string]string, 50)

	carta := map[string][]int{
		"1":    {1, 2},
		"2":    {4, 132},
		"13":   {1, 2},
		"22":   {4, 132},
		"121":  {1, 2},
		"2123": {4, 132},
	}

	for _, i := range carta {
		fmt.Println(i)
	}

	fmt.Println(carta["1"][1])

	for key, value := range carta {
		fmt.Println(key, value)
	}
}
