package main

import "fmt"

func main() {
	a := 1
	b := 1
	result := a == b

	fmt.Println(result)

	result = a != b
	fmt.Println(result)

	b = 2
	result = a == b
	fmt.Println(result)

	result = a < b
	fmt.Println(result)

	result = a <= b
	fmt.Println(result)

	b = 1
	result = a > b
	fmt.Println(result)

	result = a >= b
	fmt.Println(result)

}
