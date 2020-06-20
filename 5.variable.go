package main

import "fmt"

var x string = "hello world"

func main() {
	fmt.Println(x)
	x = "hello golang"
	fmt.Println(x)

	var y string
	fmt.Println(y)
	y = "GO"
	fmt.Println(y)

	z := "Belajar Golang"
	fmt.Println(z)

	angka := 10
	fmt.Println(angka)

	const fixed string = "tidak dapat diubah"
	fmt.Println(fixed)
	// fixed = "dapat diubah"
	fmt.Println(fixed)
}

func notMain() {
	fmt.Println(x)
}
