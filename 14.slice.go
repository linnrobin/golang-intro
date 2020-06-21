package main

import "fmt"

func main() {
	names := [5]string{
		"Andi",
		"Budi",
		"Carli",
		"Dewi",
		"Edi",
	}

	slice1 := names[1:4]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	names[1] = "not budi"

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice1[1] = "XXXXXXXXXX"

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

}
