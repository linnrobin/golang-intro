package main

import "fmt"

func main() {
	slice1 := make([]string, 3)
	slice1[0] = "Andi"
	slice1[1] = "Budi"
	slice1[2] = "Carli"

	fmt.Println(slice1)

	slice2 := append(slice1, "dewi", "edi")

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "notAndi"

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string, 10)
	copy(slice3, slice1)

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice1[0] = "veryNotAndi"

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice4 := make([]string, 2)
	copy(slice4, slice1)

	fmt.Println(slice1)
	fmt.Println(slice4)
}
