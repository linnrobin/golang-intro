package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Andi"
	names[1] = "Budi"
	names[2] = "Carli"
	fmt.Println(names)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("index", index, "=", value)
	}

	classes := [5]string{
		"Kelas 1",
		"Kelas 2",
		"Kelas 3",
		"Kelas 4",
		"Kelas 5",
	}

	for _, value := range classes {
		fmt.Println(value)
	}

}
