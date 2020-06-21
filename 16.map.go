package main

import "fmt"

func main() {
	names := make(map[string]string)
	names["1001"] = "Andi"
	names["1002"] = "Budi"
	names["1003"] = "Carli"

	fmt.Println(names)

	fmt.Println(names["1001"])
	fmt.Println(names["1002"])
	fmt.Println(names["1003"])
	fmt.Println(names["1004"])

	for NIM, name := range names {
		fmt.Println("NIM", NIM, "bernama", name)
	}
}
